from PySide6.QtCore import Qt, QTimer, QPropertyAnimation
from PySide6.QtWidgets import QLabel, QGraphicsOpacityEffect

class Notification(QLabel):
    def __init__(self, parent, text, is_error=False):
        super().__init__(parent)
        self.setText(text)
        self.setWordWrap(True)
        self.setFixedSize(220, 40)
        self.setAlignment(Qt.AlignCenter)
        
        text_color = "red" if is_error else "black"
        self.setStyleSheet(f"""
            background-color: white;
            color: {text_color};
            border-radius: 8px;
            padding: 10px;
            font-family: 'Noto Sans', sans-serif;
            font-size: 10px;
            border: 1px solid #ddd;
        """)
        
        self.opacity_effect = QGraphicsOpacityEffect(self)
        self.setGraphicsEffect(self.opacity_effect)
        
        self.adjust_position()
        
        self.anim = QPropertyAnimation(self.opacity_effect, b"opacity")
        self.anim.setDuration(300)
        self.anim.setStartValue(0.0)
        self.anim.setEndValue(1.0)
        
        self.show()
        self.anim.start()

        QTimer.singleShot(2000, self.fade_out)

    def adjust_position(self):
        if self.parent():
            p_rect = self.parent().rect()
            self.move(p_rect.width() - self.width() - 10, 10)

    def fade_out(self):
        self.anim_out = QPropertyAnimation(self.opacity_effect, b"opacity")
        self.anim_out.setDuration(600)
        self.anim_out.setStartValue(1.0)
        self.anim_out.setEndValue(0.0)
        self.anim_out.finished.connect(self.deleteLater)
        self.anim_out.start()

