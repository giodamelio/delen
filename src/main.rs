#[macro_use]
extern crate rocket;

use rocket::{
    http::Accept,
    serde::{json::Json, Serialize},
};
use rocket_dyn_templates::Template;

mod fairings;

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
        .attach(fairings::ExtensionRewrite::new(".json", Accept::JSON))
        .attach(fairings::ExtensionRewrite::new(".html", Accept::HTML))
        .mount("/", routes![index_html, index_json])
}
