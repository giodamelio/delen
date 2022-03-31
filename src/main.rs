#[macro_use]
extern crate rocket;
#[macro_use]
extern crate diesel;

use diesel::prelude::*;
use diesel::sqlite::SqliteConnection;
use rocket::{
    http::Accept,
    serde::{json::Json, Serialize},
};
use rocket_dyn_templates::Template;

use self::models::*;

mod fairings;
mod models;
mod schema;

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
    use self::schema::item::dsl::*;

    // Test the database out
    let conn = SqliteConnection::establish("db.sqlite3").expect("Cannot connect to db");

    // Add an Item
    let new_item = NewItem {
        name: "Test".into(),
        mime_type: "text/plain".into(),
        contents: vec![],
    };
    let inserted_count = diesel::insert_into(schema::item::table)
        .values(&new_item)
        .execute(&conn)
        .expect("Error inserting item");
    println!("{} items inserted", inserted_count);

    // List all the items
    let items = item.load::<Item>(&conn).expect("Cannot load items");
    println!("{:?}", items);

    rocket::build()
        .attach(Template::fairing())
        .attach(fairings::ExtensionRewrite::new(".json", Accept::JSON))
        .attach(fairings::ExtensionRewrite::new(".html", Accept::HTML))
        .mount("/", routes![index_html, index_json])
}
