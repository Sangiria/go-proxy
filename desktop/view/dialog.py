from PySide6.QtWidgets import QDialog, QComboBox, QDialogButtonBox, QLineEdit, QTextEdit, QSpinBox
from design.addsubscriptiondialog import Ui_Dialog as Ui_AddSubscription
from design.editnodedialog import Ui_Dialog as Ui_EditNode
from design.editsubdialog import Ui_Dialog as Ui_EditSubscribtion

class EditSubscribtionDialog(QDialog, Ui_EditSubscribtion):
    def __init__(self, data, parent=None):
        super().__init__(parent)
        self.setupUi(self)
        self.data = data

        self.nameEdit.textChanged.connect(self.labelError.clear)
        self.urlEdit.textChanged.connect(self.labelError.clear)

        self.nameEdit.setText(data.name)
        self.urlEdit.setText(data.url)
class EditNodeDialog(QDialog, Ui_EditNode):
    def __init__(self, data, parent=None):
        super().__init__(parent)
        self.setupUi(self)
        self.data = data
        self.start_edit = lambda: None

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

        self.fill_form()

        for widget in self.field_map.keys():
            if isinstance(widget, QLineEdit) or isinstance(widget, QLineEdit):
                widget.textChanged.connect(self.labelError.clear)
            elif isinstance(widget, QSpinBox):
                widget.valueChanged.connect(self.labelError.clear)
    def accept(self):
        delta, full = self.get_updated_data()
        
        if delta is not None:
            self.delta = delta
            self.full = full
            self.start_edit()
        else:
            if not self.labelError.text():
                super().reject()
    def fill_form(self):
        for widget, attr in self.field_map.items():
            if not self.data.HasField(attr):
                continue
            value = getattr(self.data, attr)
            if isinstance(widget, QLineEdit) or isinstance(widget, QTextEdit):
                widget.setText(str(value))
            elif isinstance(widget, QComboBox):
                index = widget.findText(str(value))
                if index >= 0:
                    widget.setCurrentIndex(index)
            elif isinstance(widget, QSpinBox):
                widget.setValue(int(value))
    def get_updated_data(self):
        delta = type(self.data)()
        delta.id = self.data.id
        if hasattr(self.data, "source_id") and self.data.source_id:
            delta.source_id = self.data.source_id

        full = type(self.data)()
        full.CopyFrom(self.data)

        required_fields = ["name", "address"]
        has_changes = False

        for widget, attr in self.field_map.items():
            if isinstance(widget, (QLineEdit, QTextEdit)):
                new_val = widget.text() if isinstance(widget, QLineEdit) else widget.toPlainText()
            elif isinstance(widget, QComboBox):
                new_val = widget.currentText()
            elif isinstance(widget, QSpinBox):
                new_val = widget.value()

            old_val = getattr(self.data, attr) if self.data.HasField(attr) else ""

            if str(new_val) != str(old_val):
                if attr in required_fields and not new_val:
                    self.labelError.setText(f"field '{attr}' is required")
                    return None, None
                
                has_changes = True
                setattr(delta, attr, new_val)
                setattr(full, attr, new_val)

        return (delta, full) if has_changes else (None, None)
    
class AddDialog(QDialog, Ui_AddSubscription):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setupUi(self)

        self.lineEdit.textChanged.connect(lambda: self.labelError.clear())