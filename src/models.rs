use super::schema::item;

#[derive(Queryable, Debug)]
pub struct Item {
    pub id: i32,
    pub name: String,
    pub mime_type: String,
    pub contents: Vec<u8>,
}

#[derive(Insertable, Debug)]
#[table_name = "item"]
pub struct NewItem {
    pub name: String,
    pub mime_type: String,
    pub contents: Vec<u8>,
}
