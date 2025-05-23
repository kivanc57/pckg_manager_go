import os
from zipfile import ZipFile
from .file_utils import join_paths, remove_file

def rename_package(src_pckg, target_pck):
    try:
        os.rename(src_pckg, target_pck)
        return target_pck
    except FileNotFoundError:
        print(f"Error: File not found: {src_pckg}")
        return None
    except Exception as e:
        print(f"An error occurred during renaming: {e}")
        return None

def extract_relevant_files_from_zip(zip_path, output_folder):
    is_placeholder_present = False

    try:
        with ZipFile(zip_path, 'r') as zObject:
            for file in zObject.namelist():
                if file.endswith(".nuspec") or file.endswith(".ps1"):
                    zObject.extract(file, output_folder)
                elif (file.endswith(".exe") or file.endswith(".msi")) and not is_placeholder_present:
                    is_placeholder_present = True
                    placeholder_path = join_paths(output_folder, "placeholder")
                    with open(placeholder_path, "w") as pHolder:
                        pHolder.write("Installer found")
        remove_file(zip_path)
        print(f"Extraction completed: {zip_path}")
    except FileNotFoundError:
        print(f"Error: File not found: {zip_path}")
    except Exception as e:
        print(f"An error occurred during extraction: {e}")
