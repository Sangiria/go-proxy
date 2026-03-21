from PySide6.QtWidgets import QDialog, QComboBox, QLineEdit, QTextEdit, QSpinBox
from design.addsubscriptiondialog import Ui_Dialog as Ui_AddSubscription
from design.editnodedialog import Ui_Dialog as Ui_EditNode
from design.editsubdialog import Ui_Dialog as Ui_EditSubscribtion
import json

class EditDialog(QDialog):
    def __init__(self, data, parent=None):
        super().__init__(parent)
        self.data = data
        self.field_map = {}
        self.start_edit = lambda: None
        self.delta = None
    def setup(self):
        self.fill_form()
        if hasattr(self, 'buttonBox'):
            self.buttonBox.accepted.disconnect()
            self.buttonBox.accepted.connect(self.handle_accept)
            
        for widget in self.field_map.keys():
            if isinstance(widget, QLineEdit) or isinstance(widget, QTextEdit):
                widget.textChanged.connect(self.labelError.clear)
    def fill_form(self):
        for widget, attr in self.field_map.items():
            if not self.data.HasField(attr):
                continue
            value = getattr(self.data, attr)

            if attr == "extra" and isinstance(widget, QTextEdit):
                try:
                    obj = json.loads(value)
                    widget.setPlainText(json.dumps(obj, indent=4, ensure_ascii=False))
                except:
                    widget.setPlainText(str(value))
            elif isinstance(widget, (QLineEdit, QTextEdit)):
                if isinstance(widget, QLineEdit): widget.setText(str(value))
                else: widget.setPlainText(str(value))
            elif isinstance(widget, QComboBox):
                index = widget.findText(str(value))
                if index >= 0: widget.setCurrentIndex(index)
            elif isinstance(widget, QSpinBox):
                widget.setValue(int(value))
    def get_updated_data(self):
        delta = type(self.data)()
        delta.id = self.data.id
        if hasattr(self.data, "source_id") and self.data.source_id:
            delta.source_id = self.data.source_id

        has_changes = False

        for widget, attr in self.field_map.items():
            if isinstance(widget, (QLineEdit, QTextEdit)):
                new_val = widget.text() if isinstance(widget, QLineEdit) else widget.toPlainText()
            elif isinstance(widget, QComboBox):
                new_val = widget.currentText()
            elif isinstance(widget, QSpinBox):
                new_val = widget.value()

            old_val = getattr(self.data, attr)

            if str(new_val) != str(old_val):
                has_changes = True
                setattr(delta, attr, new_val)

        return delta if has_changes else None
    def handle_accept(self):
        self.delta = self.get_updated_data()
        if self.delta:
            self.start_edit()
        else:
            if not self.labelError.text():
                self.reject() 

    def accept(self):
        super().accept()

class EditSubscribtionDialog(EditDialog, Ui_EditSubscribtion):
    def __init__(self, data, parent=None):
        super().__init__(data, parent)
        self.setupUi(self)
        
        self.field_map = {
            self.nameEdit: "name",
            self.urlEdit: "url"
        }
        self.setup()

class EditNodeDialog(EditDialog, Ui_EditNode):
    def __init__(self, data, parent=None):
        super().__init__(data, parent)
        self.setupUi(self)
        self.extraEdit.setStyleSheet("font-family: 'Courier New'; font-size: 10pt;")

        self.field_map = {
            self.nameEdit: "name",
            self.addressEdit: "address",
            self.portSBox: "port",
            self.uuidEdit: "uuid",
            self.transportCBox: "transport",
            self.modeCBox: "mode",
            self.securityCBox: "security",
            self.sniEdit: "sni",
            self.fpCBox: "fp",
            self.pbkEdit: "pbk",
            self.extraEdit: "extra"
        }
        self.setup()
    
class AddDialog(QDialog, Ui_AddSubscription):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setupUi(self)

        self.lineEdit.textChanged.connect(lambda: self.labelError.clear())