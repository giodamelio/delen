#[macro_use]
extern crate rocket;

use rocket::serde::{json::Json, Serialize};
use rocket_dyn_templates::Template;

#[derive(Serialize)]
#[serde(crate = "rocket::serde")]
struct Index {
    message: String,
}

#[get("/", format = "html")]
fn index_html() -> Template {
    let context = Index {
        message: "Hello World!".into(),
    };

    Template::render("index", context)
}

#[get("/", format = "json", rank = 2)]
fn index_json() -> Json<Index> {
    let context = Index {
        message: "Hello World!".into(),
    };

    Json(context)
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .attach(Template::fairing())
        .mount("/", routes![index_html, index_json])
}
