import pandas as pd
import isodate

LAST_DAY = 7
BEST_ALPHA = 0.37926901907322497


def transform_category_feature(df):
    feature = round(df.groupby("category_id").size() / len(df), 2)
    df.loc[:, "category_ratio"] = df["category_id"].map(feature)


regions = [
    "FI", "KZ", "UA", "EG", "IT", "PT",
    "GB", "BE", "NO", "LV", "RU", "TR",
    "ES", "FR", "SE", "PL", "NL", "DE"]


def input_cleaning(region):
    input_path = "data_" + region + ".tsv"
    data = pd.read_csv(input_path, delimiter="\t")
    data = data.replace(to_replace="None", value=0)
    for item in data.columns[5:]:
        data[item] = pd.to_numeric(data[item])
    data["duration_in_seconds"] = data.duration.map(lambda x: isodate.parse_duration(x).total_seconds())
    data["size"] = data.duration_in_seconds.map(lambda x: round(x / 60 * 30, 2))
    transform_category_feature(data)
    data.reset_index(drop=True, inplace=True)

    return data

def calc_wma(column, days, df):
    res = 0
    i = days
    s = (1 + days) / 2 * days
    while i > 0:
        res += df[column + '_' + str(i)] * i
        i -= 1
        
    return res / s


def update_df(df):
    df["views"] = df["views_" + str(LAST_DAY)]
    df["comments"] = df["comments_" + str(LAST_DAY)]
    df["likes"] = df["likes_" + str(LAST_DAY)]
    
    df["views_wma"] = df.apply(lambda x: calc_wma("views", LAST_DAY, x), axis=1)
    df["comments_wma"] = df.apply(lambda x: calc_wma("comments", LAST_DAY, x), axis=1)
    df["likes_wma"] = df.apply(lambda x: calc_wma("likes", LAST_DAY, x), axis=1)

    return df


def output(region, df):
    output_columns = [
        "video_id", "title", "size", "category_ratio", 
        "views", "comments", "likes",
        "views_wma", "comments_wma", "likes_wma",
    ]
    output = df[output_columns]
    output_path_csv = region + "/cleaned_data.tsv"
    output.to_csv(output_path_csv, sep='\t', index=False)
    output_path_json = region + "/video_set.json"
    out = output.to_json(orient='records')
    with open(output_path_json, 'w') as f:
        f.write(out)
        

def main():
    for region in regions:
        # input & cleaning
        data = input_cleaning(region)

        # update df
        data = update_df(data)

        # output
        output(region, data)


if __name__ == "__main__":
    main()
