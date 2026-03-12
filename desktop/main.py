import sys
from PyQt6.QtWidgets import QApplication
from logic.window import MainWindow
from service.worker import channel

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()

    exit_code = app.exec()
    channel.close()

    sys.exit(exit_code)