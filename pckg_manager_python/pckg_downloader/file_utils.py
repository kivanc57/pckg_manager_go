import os

def join_paths(*args):
    return os.path.join(*args)

def remove_file(path):
    try:
        if os.path.exists(path):
            os.remove(path)
    except Exception as e:
        print(f"Error removing file: {e}")

def create_directory(target_dir):
    try:
        os.makedirs(target_dir, 0o755, exist_ok=True)
    except OSError as e:
        print(f"Directory {target_dir} can not be created: {e}")
