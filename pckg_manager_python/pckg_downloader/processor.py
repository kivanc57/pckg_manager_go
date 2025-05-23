from .file_utils import join_paths, create_directory
from .download_utils import download_file
from .package_utils import rename_package, extract_relevant_files_from_zip

def process_package(pckg_name, pckg_ver, links):
    """Process one package: download, rename and extract."""
    for link in links:
        output_file = f"{pckg_name}.{pckg_ver}.nupkg"
        output_folders = join_paths("output", pckg_name, pckg_ver)
        output_path = join_paths(output_folders, output_file)
        zip_path = output_path.replace(".nupkg", ".zip")

        if os.path.exists(zip_path):
            print(f"Already processed: {zip_path}")
            return False

        create_directory(output_folders)
        pckg_link = f"{link}&path={output_file}"

        if download_file(pckg_link, output_path):
            new_path = rename_package(output_path, zip_path)
            if new_path:
                extract_relevant_files_from_zip(new_path, output_folders)
                print("------------------------------------------------------------------")
                return True
    return False
