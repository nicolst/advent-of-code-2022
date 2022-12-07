from dataclasses import dataclass


@dataclass
class File:
    name: str
    size: int


class Folder:
    def __init__(self, parent: Folder | None):
        self.parent = parent
