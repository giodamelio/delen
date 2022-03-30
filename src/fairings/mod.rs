use rocket::{
    fairing::{Fairing, Info, Kind},
    http::{Accept, Header},
    Data, Request,
};

pub struct ExtensionRewrite {
    extension: &'static str,
    header: Header<'static>,
}

impl ExtensionRewrite {
    pub fn new(extension: &'static str, header: Accept) -> Self {
        ExtensionRewrite {
            extension,
            header: header.into(),
        }
    }
}

#[rocket::async_trait]
impl Fairing for ExtensionRewrite {
    fn info(&self) -> Info {
        Info {
            name: "Extension Rewriter",
            kind: Kind::Request,
        }
    }

    async fn on_request(&self, request: &mut Request<'_>, _data: &mut Data<'_>) {
        let origin = request.uri().clone();

        // If the path ends with a file extension, remove it and add a matching Accept header
        // If the mapping breaks, keep the original path
        let new_origin = if origin.path().ends_with(self.extension) {
            // Add accept header
            request.replace_header(self.header.clone());

            // Remove the extension from the path
            // It is safe to unwrap the `strip_suffix` because we already checked it's existance
            origin
                .map_path(|path| path.strip_suffix(self.extension).unwrap())
                .unwrap_or_else(|| origin.to_owned())
        } else {
            origin.to_owned()
        };

        request.set_uri(new_origin);
    }
}
