APP_NAME = ProxyApp
BUILD_DIR = $(CURDIR)/dist/$(APP_NAME)
UI_ENTRY = desktop/main.py
BE_ENTRY = core/main.go
XRAY_SOURCE = bin/xray

.PHONY: all clean build_ui build_be copy_assets fix_perms init

all: clean init build_be build_ui copy_assets fix_perms
	@echo "-------------------------------------------------------"
	@echo "Сборка завершена успешно!"
	@echo "Все файлы приложения находятся здесь: $(BUILD_DIR)"
	@echo "-------------------------------------------------------"

init:
	mkdir -p $(BUILD_DIR)/bin
	mkdir -p $(BUILD_DIR)/design/form

build_be:
	@echo ">>> Сборка Go-бэкенда..."
	cd core && go mod tidy && go build -o $(BUILD_DIR)/core_backend main.go

build_ui:
	@echo ">>> Сборка Python UI через PyInstaller..."
	pyinstaller --noconfirm --onedir --windowed \
		--contents-directory .lib \
		--add-data "desktop/design/form/icon.png:design/form" \
		--distpath ./dist \
		--workpath ./build_tmp \
		--name $(APP_NAME)_UI $(UI_ENTRY)
	
	@echo ">>> Перенос UI в финальную директорию..."
	rsync -a ./dist/$(APP_NAME)_UI/ $(BUILD_DIR)/
	rm -rf ./dist/$(APP_NAME)_UI ./build_tmp *.spec

copy_assets:
	@echo ">>> Копирование Xray и создание скрипта запуска..."
	cp $(XRAY_SOURCE) $(BUILD_DIR)/bin/xray
	
	@echo '#!/bin/bash' > $(BUILD_DIR)/run.sh
	@echo 'cd "$$(dirname "$$0")"' >> $(BUILD_DIR)/run.sh
	@echo '# Запуск бэкенда' >> $(BUILD_DIR)/run.sh
	@echo './core_backend &' >> $(BUILD_DIR)/run.sh
	@echo 'BE_PID=$$!' >> $(BUILD_DIR)/run.sh
	@echo '# Запуск интерфейса' >> $(BUILD_DIR)/run.sh
	@echo './$(APP_NAME)_UI' >> $(BUILD_DIR)/run.sh
	@echo '# Завершение бэкенда' >> $(BUILD_DIR)/run.sh
	@echo 'kill -SIGTERM $$BE_PID' >> $(BUILD_DIR)/run.sh
	@echo 'wait $$BE_PID' >> $(BUILD_DIR)/run.sh

fix_perms:
	@echo ">>> Настройка прав доступа..."
	chmod +x $(BUILD_DIR)/core_backend
	chmod +x $(BUILD_DIR)/bin/xray
	chmod +x $(BUILD_DIR)/$(APP_NAME)_UI
	chmod +x $(BUILD_DIR)/run.sh

clean:
	@echo ">>> Полная очистка временных файлов..."
	rm -rf dist build build_tmp *.spec