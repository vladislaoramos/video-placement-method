import datetime

from scrapper.handler import DataHandler
from scrapper.parser import DataParser
from scrapper.reader import Reader, get_video_set, get_country_codes
from scrapper.logger import get_logger
from scrapper.setups.yt import API_KEYS_COUNT


logger = get_logger()


def main():
    logger.debug("RUN SCRAPPER")

    start_time = datetime.datetime.now()

    countries = get_country_codes()
    logger.debug(f"getting video set from trends, number of regions: {len(countries)}")

    for country in countries:
        video_set = get_video_set(country)

        dp = DataParser(video_set, API_KEYS_COUNT)
        received_data = dp.get_response()

        dr = Reader(country)
        records = dr.get_records()

        is_first_day = len(records.get("views")) == 0
        dh = DataHandler(received_data, records, is_first_day, country)
        dh.save_data()

    logger.debug("FINISH SCRAPPER")

    finish_time = datetime.datetime.now()
    logger.debug("Time: " + str(finish_time - start_time))


if __name__ == "__main__":
    main()
    # test()

# python start.py --days 10 --key_path scrapper/setups/api-key.txt --country_code_path scrapper/setups/country-codes.txt
# python start.py --days 7 --key_path scrapper/setups/api-key.txt --country_code_path scrapper/setups/country-codes.txt
# python start.py --days 5 --key_path scrapper/setups/api-key.txt --country_code_path scrapper/setups/country-codes.txt
# python start.py --days 14 --key_path scrapper/setups/api-key.txt --country_code_path scrapper/setups/country-codes.txt
# python start.py
