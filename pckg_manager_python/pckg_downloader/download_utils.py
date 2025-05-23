import requests

def download_file(url, output_path):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            with open(output_path, "wb") as file:
                file.write(response.content)
            print(f"Download succeeded from: {url}")
            print(f"File is written to: {output_path}")
            return True
        else:
            print(f"Download failed from: {url} with status code {response.status_code}")
            print("------------------------------------------------------------------")
            return False
    except Exception as e:
        print(f"Error downloading file: {e}")
        print("------------------------------------------------------------------")
        return False
