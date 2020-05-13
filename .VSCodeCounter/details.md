# Details

Date : 2020-05-13 14:18:17

Directory /mnt/b/Development/GitHub/lambdish-compiler

Total : 54 files,  15851 codes, 0 comments, 1064 blanks, all 16915 lines

[summary](results.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [ast/ast.go](/ast/ast.go) | Go | 290 | 0 | 69 | 359 |
| [ast/astx.go](/ast/astx.go) | Go | 278 | 0 | 75 | 353 |
| [ast/astx_test.go](/ast/astx_test.go) | Go | 6 | 0 | 2 | 8 |
| [cmd/clamb/clamb.go](/cmd/clamb/clamb.go) | Go | 73 | 0 | 24 | 97 |
| [cmd/rlamb/rlamb.go](/cmd/rlamb/rlamb.go) | Go | 33 | 0 | 11 | 44 |
| [dir/festack.go](/dir/festack.go) | Go | 36 | 0 | 9 | 45 |
| [dir/funcdir.go](/dir/funcdir.go) | Go | 107 | 0 | 28 | 135 |
| [dir/funcdir_test.go](/dir/funcdir_test.go) | Go | 780 | 0 | 7 | 787 |
| [dir/vardir.go](/dir/vardir.go) | Go | 60 | 0 | 20 | 80 |
| [dir/vardir_test.go](/dir/vardir_test.go) | Go | 290 | 0 | 9 | 299 |
| [gocc/errors/errors.go](/gocc/errors/errors.go) | Go | 49 | 0 | 8 | 57 |
| [gocc/lexer/acttab.go](/gocc/lexer/acttab.go) | Go | 204 | 0 | 8 | 212 |
| [gocc/lexer/lexer.go](/gocc/lexer/lexer.go) | Go | 157 | 0 | 17 | 174 |
| [gocc/lexer/transitiontable.go](/gocc/lexer/transitiontable.go) | Go | 600 | 0 | 4 | 604 |
| [gocc/parser/action.go](/gocc/parser/action.go) | Go | 42 | 0 | 10 | 52 |
| [gocc/parser/actiontable.go](/gocc/parser/actiontable.go) | Go | 5,419 | 0 | 4 | 5,423 |
| [gocc/parser/gototable.go](/gocc/parser/gototable.go) | Go | 3,129 | 0 | 5 | 3,134 |
| [gocc/parser/parser.go](/gocc/parser/parser.go) | Go | 184 | 0 | 33 | 217 |
| [gocc/parser/productionstable.go](/gocc/parser/productionstable.go) | Go | 352 | 0 | 5 | 357 |
| [gocc/token/token.go](/gocc/token/token.go) | Go | 146 | 0 | 23 | 169 |
| [gocc/util/litconv.go](/gocc/util/litconv.go) | Go | 97 | 0 | 12 | 109 |
| [gocc/util/rune.go](/gocc/util/rune.go) | Go | 36 | 0 | 4 | 40 |
| [grammar/grammar_test.go](/grammar/grammar_test.go) | Go | 44 | 0 | 14 | 58 |
| [ic/addressstack.go](/ic/addressstack.go) | Go | 63 | 0 | 17 | 80 |
| [ic/generateaddresses.go](/ic/generateaddresses.go) | Go | 97 | 0 | 20 | 117 |
| [ic/generateaddresses_test.go](/ic/generateaddresses_test.go) | Go | 41 | 0 | 10 | 51 |
| [ic/generatecode.go](/ic/generatecode.go) | Go | 468 | 0 | 116 | 584 |
| [ic/generator.go](/ic/generator.go) | Go | 133 | 0 | 34 | 167 |
| [ic/ic.go](/ic/ic.go) | Go | 47 | 0 | 11 | 58 |
| [ic/ic_test.go](/ic/ic_test.go) | Go | 58 | 0 | 17 | 75 |
| [ic/util.go](/ic/util.go) | Go | 20 | 0 | 5 | 25 |
| [integration/integration_test.go](/integration/integration_test.go) | Go | 65 | 0 | 19 | 84 |
| [mem/mem.go](/mem/mem.go) | Go | 334 | 0 | 42 | 376 |
| [mem/util.go](/mem/util.go) | Go | 5 | 0 | 2 | 7 |
| [quad/quad.go](/quad/quad.go) | Go | 231 | 0 | 22 | 253 |
| [sem/funccheck.go](/sem/funccheck.go) | Go | 109 | 0 | 25 | 134 |
| [sem/funccheck_test.go](/sem/funccheck_test.go) | Go | 36 | 0 | 10 | 46 |
| [sem/funcutil.go](/sem/funcutil.go) | Go | 15 | 0 | 5 | 20 |
| [sem/scopecheck.go](/sem/scopecheck.go) | Go | 99 | 0 | 18 | 117 |
| [sem/scopecheck_test.go](/sem/scopecheck_test.go) | Go | 41 | 0 | 11 | 52 |
| [sem/scopeutil.go](/sem/scopeutil.go) | Go | 23 | 0 | 6 | 29 |
| [sem/sem.go](/sem/sem.go) | Go | 45 | 0 | 8 | 53 |
| [sem/sem_test.go](/sem/sem_test.go) | Go | 53 | 0 | 14 | 67 |
| [sem/semanticcube.go](/sem/semanticcube.go) | Go | 230 | 0 | 34 | 264 |
| [sem/typecheck.go](/sem/typecheck.go) | Go | 100 | 0 | 25 | 125 |
| [sem/typecheck_test.go](/sem/typecheck_test.go) | Go | 45 | 0 | 13 | 58 |
| [sem/typeutil.go](/sem/typeutil.go) | Go | 167 | 0 | 30 | 197 |
| [types/types.go](/types/types.go) | Go | 127 | 0 | 25 | 152 |
| [types/types_test.go](/types/types_test.go) | Go | 172 | 0 | 8 | 180 |
| [vm/mem.go](/vm/mem.go) | Go | 282 | 0 | 22 | 304 |
| [vm/operations.go](/vm/operations.go) | Go | 81 | 0 | 27 | 108 |
| [vm/util.go](/vm/util.go) | Go | 51 | 0 | 17 | 68 |
| [vm/vm.go](/vm/vm.go) | Go | 166 | 0 | 39 | 205 |
| [vm/vm_test.go](/vm/vm_test.go) | Go | 35 | 0 | 11 | 46 |

[summary](results.md)