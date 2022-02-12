use mongodb::{bson::{doc}, Client};
use futures_util::TryStreamExt;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct Trade {
    tp:  f64,
    tv:  f64,
    ab:  String,
    tms: u64,
}

#[tokio::main]
async fn main() -> mongodb::error::Result<()> {
    let client = Client::with_uri_str("mongodb://localhost:27017").await?;

    let collection = client
        .database("trades")
        .collection::<Trade>("XRP");

    let mut cursor = collection.find(None, None).await?;

    println!("trade_price,trade_volume,ask_bid,timestamp");
    while let Some(trade) = cursor.try_next().await? {
        println!("{},{},{},{}", trade.tp, trade.tv, trade.ab, trade.tms);
    }

    Ok(())
}
