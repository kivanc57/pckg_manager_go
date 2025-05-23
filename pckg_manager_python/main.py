import time
import os
from dotenv import load_dotenv

from pckg_downloader.parser import parse_package_line
from pckg_downloader.processor import process_package

load_dotenv()

def main():
    input_file = os.path.join("data", "final_list.txt")
    links = os.getenv("LINKS").split(",")
    count = 0

    try:
        with open(input_file, "r") as f:
            for line in f:
                pckg_name, pckg_ver = parse_package_line(line)
                if not pckg_name or not pckg_ver:
                    continue
                if process_package(pckg_name, pckg_ver, links):
                    count += 1
    except FileNotFoundError:
        print(f"Input file not found: {input_file}")
    except Exception as e:
        print(f"An error occurred reading input file: {e}")

    print(f"PACKAGE COUNT: {count}")

if __name__ == "__main__":
    start = time.time()
    main()
    end = time.time()
    print(f"Program took {end - start} seconds...")
