// Below is a hack to make Rocket and Anyhow play nice
// Stolen from https://crates.io/crates/rocket_anyhow
pub type Result<T = ()> = std::result::Result<T, Error>;

#[derive(Debug)]
pub struct Error(pub anyhow::Error);

impl<E> From<E> for Error
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
