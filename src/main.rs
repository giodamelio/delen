use std::collections::HashMap;
use std::path::Path;

use rocket::http::Accept;
use rocket::serde::json::Json;
use rocket::serde::Serialize;
use rocket::tokio::time::{sleep, Duration};
use rocket::{get, routes};
use rocket_dyn_templates::Template;
use sqlx::{migrate::Migrator, sqlite, Pool, Sqlite};

mod fairings;

#[derive(Serialize, Debug)]
#[serde(crate = "rocket::serde")]
struct File {
    name: &'static str,
}

// Type alias for the db connection
type DB = Pool<Sqlite>;

#[get("/", format = "html", rank = 1)]
async fn index_html(pool: &rocket::State<DB>) -> Result<Template> {
    let files: Vec<File> = vec![File { name: "foo.txt" }, File { name: "bar.txt" }];
    let mut map: HashMap<&str, Vec<File>> = HashMap::new();
    map.insert("files", files);

    let row = sqlx::query!("SELECT (150) as num")
        .fetch_one(pool.inner())
        .await?;
    println!("Result: {:?}", row);

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

// Below is a hack to make Rocket and Anyhow play nice
// Stolen from https://crates.io/crates/rocket_anyhow
pub type Result<T = ()> = std::result::Result<T, Error>;

#[derive(Debug)]
pub struct Error(pub anyhow::Error);

impl<E> From<E> for crate::Error
where
    E: Into<anyhow::Error>,
{
    fn from(error: E) -> Self {
        Error(error.into())
    }
}

impl<'r, 'o: 'r> rocket::response::Responder<'r, 'o> for Error {
    fn respond_to(self, request: &rocket::Request<'_>) -> rocket::response::Result<'o> {
        rocket::response::Debug(self.0).respond_to(request)
    }
}
// </hack>
