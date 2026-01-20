@echo off
chcp 65001 >nul
REM === Переименование директорий с кириллицы в латиницу ===

ren "алгоритмы на строках" string_algorithms
ren "бинарный поиск" binary_search
ren "деревья поиска" tree_search
ren "динамическое программирование" dynamic_programming
ren "линейный поиск" linear_search
ren "массивы" arrays
ren "очереди, деки и стек" queues_stacks
ren "саморасширяющиеся массивы" dynamic_arrays
ren "сортировки" sorting
ren "сортировки\продвинутые сортировки" advanced_sorting
ren "списки" lists
ren "хэш-таблица" hash_table
ren "экспоненциальный и тернарный поиск" exponential_ternary_search

echo.
echo === Готово! Все папки переименованы ===
pause