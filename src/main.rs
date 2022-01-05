use std::collections::HashMap;

use rocket::http::Accept;
use rocket::serde::json::Json;
use rocket::serde::Serialize;
use rocket::tokio::time::{sleep, Duration};
use rocket::{get, routes};
use rocket_dyn_templates::Template;

mod fairings;

#[derive(Serialize, Debug)]
#[serde(crate = "rocket::serde")]
struct File {
    name: &'static str,
}

#[get("/", format = "html", rank = 1)]
fn index_html() -> Template {
    let files: Vec<File> = vec![File { name: "foo.txt" }, File { name: "bar.txt" }];
    let mut map: HashMap<&str, Vec<File>> = HashMap::new();
    map.insert("files", files);

    Template::render("index", map)
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

// #[launch]
// fn rocket() -> _ {
//     rocket::build()
//         .mount("/", routes![index_json, index_html, hello, delay])
//         .attach(Template::fairing())
//         .attach(fairings::ExtensionRewrite::new(".json", Accept::JSON))
//         .attach(fairings::ExtensionRewrite::new(".html", Accept::HTML))
// }

#[tokio::main]
async fn main() -> Result<(), rocket::Error> {
    rocket::build()
        .mount("/", routes![index_json, index_html, hello, delay])
        .attach(Template::fairing())
        .attach(fairings::ExtensionRewrite::new(".json", Accept::JSON))
        .attach(fairings::ExtensionRewrite::new(".html", Accept::HTML))
        .ignite()
        .await?
        .launch()
        .await
}
