import os

import pandas as pd
from dotenv import load_dotenv
from fastapi import FastAPI
from pydantic import BaseModel

load_dotenv()

app = FastAPI(title="Reports API")


class Sale(BaseModel):
    sku: str
    quantity: int
    unit_price: float


@app.get("/health")
def health():
    return {"status": "ok", "env": os.getenv("APP_ENV", "dev")}


@app.post("/summary")
def summary(sales: list[Sale]):
    df = pd.DataFrame([s.model_dump() for s in sales])
    df["total"] = df["quantity"] * df["unit_price"]
    return {"rows": len(df), "revenue": float(df["total"].sum())}
