import pandas as pd
from sklearn.compose import ColumnTransformer
from sklearn.pipeline import Pipeline
from sklearn.preprocessing import OneHotEncoder, StandardScaler

from sklearn.base import BaseEstimator, TransformerMixin

poi_list = [
    "School",
    "Clinic",
    "PostOffice",
    "Kindergarten",
    "Restaurant",
    "College",
    "Pharmacy",
]

numeric_cols = [
    "SquareMeters",
    "Rooms",
    "Floor",
    "FloorCount",
    "CentreDistance",
    "PoiCount",
    "Age",
]

orig_bool_cols = [
    "HasParkingSpace",
    "HasBalcony",
    "HasElevator",
    "HasSecurity",
    "HasStorageRoom",
]
bool_cols = orig_bool_cols + [f"Has{poi}" for poi in poi_list]

cat_cols = ["City", "Type"]


class Preprocessor(BaseEstimator, TransformerMixin):
    poi_list = poi_list
    bool_cols = orig_bool_cols
    drop_cols = ["Latitude", "Longitude", "Ownership"]

    def fit(self, X, y=None):
        return self

    def transform(self, X):
        X = X.copy()
        X = X.drop(columns=self.drop_cols)

        for poi in self.poi_list:
            dist_col = poi + "Distance"
            flag_col = "Has" + poi
            X[flag_col] = X[dist_col].apply(lambda x: 1 if x <= 1 else 0)
        X = X.drop(columns=[poi + "Distance" for poi in self.poi_list])

        for col in self.bool_cols:
            X[col] = X[col].apply(lambda x: 1 if x == "yes" else 0)

        X["Age"] = 2025 - X["BuildYear"]
        X = X.drop(columns=["BuildYear"])

        return X


pipeline = Pipeline(
    [
        ("custom", Preprocessor()),
        (
            "cols",
            ColumnTransformer(
                [
                    ("num", Pipeline([("scale", StandardScaler())]), numeric_cols),
                    ("bool", "passthrough", bool_cols),
                    ("cat", Pipeline([("ohe", OneHotEncoder())]), cat_cols),
                ]
            ),
        ),
    ]
)


def transform(rows):
    df = pd.DataFrame(rows)

    df = df.drop(columns=["BuildingMaterial", "Condition"])
    df = df[df["Type"].str.len() > 0]
    df = df[df["HasElevator"].str.len() > 0]

    numeric_cols = df.select_dtypes("number").columns
    cols_zero_means_missing = [c for c in numeric_cols if c != "PoiCount"]

    mask_zero_miss = (df[cols_zero_means_missing] == 0).any(axis=1)
    df = df[~mask_zero_miss]

    df = df.reset_index(drop=True)

    uuid = df.pop("UUID")
    price = df.pop("Price")

    df = pd.DataFrame(
        pipeline.fit_transform(df),
        columns=pipeline.named_steps["cols"].get_feature_names_out(),
    )

    df.insert(0, "UUID", uuid)
    df["Price"] = price

    return df.to_dict(orient="records")
