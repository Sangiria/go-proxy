import os
import sys

os.environ["QT_QPA_PLATFORM"] = "xcb"

from PySide6.QtWidgets import QApplication
from view.window import MainWindow
from model.worker import channel

import faulthandler
faulthandler.enable()

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()

    exit_code = app.exec()
    channel.close()

    sys.exit(exit_code)