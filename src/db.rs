use rocket::{
    http::Status,
    request::{FromRequest, Outcome, Request},
};
use sqlx::{Pool, Sqlite};

use crate::error::Error;

// Type alias for the db connection
pub type DB = Pool<Sqlite>;

#[derive(Debug)]
pub struct ItemCount(pub i32);

#[rocket::async_trait]
impl<'r> FromRequest<'r> for ItemCount {
    type Error = Error;

    async fn from_request(request: &'r Request<'_>) -> Outcome<Self, Self::Error> {
        // Get the db connection from the managed state
        let db: &DB = match request.rocket().state::<DB>() {
            Some(d) => d,
            None => panic!("DB Connection must exist"),
        };

        let count = sqlx::query!("SELECT count(*) as count FROM item")
            .fetch_one(db)
            .await;

        match count {
            Ok(r) => Outcome::Success(ItemCount(r.count)),
            Err(err) => Outcome::Failure((Status::InternalServerError, err.into())),
        }
    }
}
