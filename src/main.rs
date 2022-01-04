use std::collections::HashMap;

use rocket::fairing::AdHoc;
use rocket::http::Accept;
use rocket::serde::json::Json;
use rocket::serde::Serialize;
use rocket::tokio::time::{sleep, Duration};
use rocket::{get, launch, routes};
use rocket_dyn_templates::Template;

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

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![index_json, index_html, hello, delay])
        .attach(Template::fairing())
        .attach(AdHoc::on_request(
            "ExtensionOverride",
            |request, _response| {
                Box::pin(async move {
                    let origin = request.uri().clone();

                    // If the path ends with a file extension, remove it and add a matching Accept header
                    // If the mapping breaks, keep the original path
                    let new_origin = if origin.path().ends_with(".json") {
                        // Add accept header
                        request.replace_header(Accept::JSON);

                        // Remove the extension from the path
                        // It is safe to unwrap the `strip_suffix` because we already checked it's existance
                        origin
                            .map_path(|path| path.strip_suffix(".json").unwrap())
                            .unwrap_or_else(|| origin.to_owned())
                    } else {
                        origin.to_owned()
                    };

                    request.set_uri(new_origin);
                })
            },
        ))
}
