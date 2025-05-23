def parse_package_line(line):
    """Extract package name and version from the input line."""
    line = line.strip()
    if not line or any(c.isspace() for c in line):
        return None, None
    try:
        i = line.index(":")
        pckg_name = line[:i]
        pckg_ver = line[i+1:].strip()
        return pckg_name, pckg_ver
    except ValueError:
        print(f"Invalid package line format: {line}")
        return None, None
