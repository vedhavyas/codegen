use actix_web::{web, App, HttpServer, Responder};
use actix_web::middleware::Logger;
use std::env;
use dotenv::dotenv;

mod db;
mod models;

async fn index() -> impl Responder {
    "Note Taking App"
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    env_logger::init();
    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");

    db::run_db(database_url);

    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .service(web::resource("/").to(index))
    })
    .bind("127.0.0.1:8000")?
    .run()
    .await
}