import sys
from PyQt6.QtWidgets import QApplication
from view.window import MainWindow
from model.worker import channel

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()

    exit_code = app.exec()
    channel.close()

    sys.exit(exit_code)