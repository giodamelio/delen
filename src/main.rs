use std::collections::HashMap;
use std::path::Path;

use rocket::http::Accept;
use rocket::serde::json::Json;
use rocket::serde::Serialize;
use rocket::tokio::time::{sleep, Duration};
use rocket::{get, routes};
use rocket_dyn_templates::Template;
use sqlx::{migrate::Migrator, sqlite};

use crate::db::item::Item;
use crate::db::DB;
use crate::error::Result;

mod db;
mod error;
mod fairings;

#[derive(Serialize, Debug)]
#[serde(crate = "rocket::serde")]
struct File {
    name: &'static str,
}

#[get("/", format = "html", rank = 1)]
async fn index_html(pool: &rocket::State<DB>, item_count: db::ItemCount) -> Result<Template> {
    let files: Vec<File> = vec![File { name: "foo.txt" }, File { name: "bar.txt" }];
    let mut map: HashMap<&str, Vec<File>> = HashMap::new();
    map.insert("files", files);

    let item = sqlx::query_as!(Item, "SELECT * FROM item")
        .fetch_all(pool.inner())
        .await;

    println!(
        "Local query result: {:?}, Guard result: {:?}",
        item, item_count
    );

    Ok(Template::render("index", map))
}

#[get("/", format = "json", rank = 2)]
fn index_json() -> Json<Vec<File>> {
    let files: Vec<File> = vec![File { name: "foo.txt" }, File { name: "bar.txt" }];
    Json(files)
}

#[get("/hello/<name>")]
fn hello(name: &str) -> String {
    format!("Hello, {}!", name)
}

#[get("/delay/<seconds>")]
async fn delay(seconds: u64) -> String {
    sleep(Duration::from_secs(seconds)).await;
    format!("Waited for {} seconds", seconds)
}

#[tokio::main]
async fn main() -> Result<()> {
    // Connect to the database
    let pool = sqlite::SqlitePoolOptions::new()
        .connect("sqlite://db.sqlite")
        .await?;

    // Run any outstanding migrations
    let migrator = Migrator::new(Path::new("./migrations")).await?;
    migrator.run(&pool).await?;

    // Configure the server
    let rocket = rocket::build()
        .mount("/", routes![index_json, index_html, hello, delay])
        .attach(Template::fairing())
        .attach(fairings::ExtensionRewrite::new(".json", Accept::JSON))
        .attach(fairings::ExtensionRewrite::new(".html", Accept::HTML))
        .manage::<DB>(pool);

    // Start the server
    rocket.ignite().await?.launch().await?;

    Ok(())
}
