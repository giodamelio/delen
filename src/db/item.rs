use sqlx::types::chrono::NaiveDateTime;

use crate::db::DB;
use crate::error::{Error, Result};

#[derive(Debug)]
pub struct Item {
    pub id: i64,
    pub filename: String,
    pub contents: Vec<u8>,
    pub filetype: Option<String>,
    pub created_at: NaiveDateTime,
}

pub async fn get_items(db: &DB) -> Result<Vec<Item>> {
    sqlx::query_as!(Item, "SELECT * FROM item")
        .fetch_all(db)
        .await
        .map_err(Error)
}
