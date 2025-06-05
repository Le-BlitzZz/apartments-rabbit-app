import trainer.config as config
import requests, time
import numpy as np
import pandas as pd

from sklearn.linear_model import LinearRegression
from sklearn.ensemble import RandomForestRegressor
from sklearn.model_selection import cross_validate
from sklearn.metrics import r2_score, root_mean_squared_error
from sklearn.model_selection import train_test_split


def wait_for_api():
    while True:
        try:
            r = requests.get(config.API, timeout=5)

            if r.status_code == 204:
                print("Server up - data not loaded yet")
            elif r.status_code == 200:
                loaded = r.headers.get("X-Loaded", "").lower() == "true"
                if loaded:
                    print("Data is loaded!")
                    return r.json()
            else:
                print(f"Unexpected status {r.status_code}")
        except requests.RequestException as e:
            print(f"API not reachable: {e}")

        print("Waiting 10 s …")
        time.sleep(10)


def prepare(df):
    X = df.drop(columns=["ID", "CreatedAt", "UpdatedAt", "DeletedAt", "UUID"])
    y = X.pop("Price")
    return X, y


def main():
    data = wait_for_api()
    df = pd.DataFrame(data)
    print("df.shape:", df.shape)

    X, y = prepare(df)

    X_temp, X_test, y_temp, y_test = train_test_split(
        X, y, test_size=0.1, random_state=42
    )

    X_train, X_val, y_train, y_val = train_test_split(
        X_temp, y_temp, test_size=0.222, random_state=42
    )  # 0.22 is ~20% of the total

    print(
        "Data prepared: ",
        X_train.shape,
        X_val.shape,
        X_test.shape,
        y_train.shape,
        y_val.shape,
        y_test.shape,
    )

    models = {
        "LinearRegression": LinearRegression(),
        "RandomForest": RandomForestRegressor(n_estimators=100, random_state=42),
    }

    cv_results = {}
    for name, model in models.items():
        print(f"Running 5-fold CV for {name}...")
        res = cross_validate(
            model,
            X_train,
            y_train,
            cv=5,
            scoring=["r2", "neg_root_mean_squared_error"],
            return_train_score=True,
        )
        cv_results[name] = {
            "train_r2": np.mean(res["train_r2"]),
            "valid_r2": np.mean(res["test_r2"]),
            "train_rmse": -np.mean(res["train_neg_root_mean_squared_error"]),
            "valid_rmse": -np.mean(res["test_neg_root_mean_squared_error"]),
        }
        print(
            f"  train squared R={cv_results[name]['train_r2']:.3f}, "
            f"valid squared R={cv_results[name]['valid_r2']:.3f}"
        )
        print(
            f"  train RMSE={cv_results[name]['train_rmse']:.0f}, "
            f"valid RMSE={cv_results[name]['valid_rmse']:.0f}"
        )

    best_name = max(cv_results, key=lambda n: cv_results[n]["valid_r2"])
    best_model = models[best_name]
    print(f"\nBest model by CV: {best_name}")

    best_model.fit(X_train, y_train)

    for split_name, (X, y) in [
        ("Validation", (X_val, y_val)),
        ("Test", (X_test, y_test)),
    ]:
        preds = best_model.predict(X)
        r2 = r2_score(y, preds)
        rmse = root_mean_squared_error(y, preds)
        print(f"{split_name} performance ({best_name}):")
        print(f"  squared R   = {r2:.3f}")
        print(f"  RMSE = {rmse:.0f}")


if __name__ == "__main__":
    main()
