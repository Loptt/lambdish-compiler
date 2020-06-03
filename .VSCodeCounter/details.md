# Details

Date : 2020-05-31 21:58:15

Directory /Users/vsapiens/Documents/lambdish-compiler

Total : 114 files,  20095 codes, 654 comments, 1668 blanks, all 22417 lines

[summary](results.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [README.md](/README.md) | Markdown | 5 | 0 | 2 | 7 |
| [ast/ast.go](/ast/ast.go) | Go | 231 | 58 | 69 | 358 |
| [ast/astx.go](/ast/astx.go) | Go | 279 | 27 | 85 | 391 |
| [cmd/clamb/clamb.go](/cmd/clamb/clamb.go) | Go | 73 | 0 | 24 | 97 |
| [cmd/rlamb/rlamb.go](/cmd/rlamb/rlamb.go) | Go | 33 | 0 | 11 | 44 |
| [cmd/tests/test4.lsh](/cmd/tests/test4.lsh) | Lambdish | 11 | 0 | 1 | 12 |
| [dir/festack.go](/dir/festack.go) | Go | 30 | 6 | 9 | 45 |
| [dir/funcdir.go](/dir/funcdir.go) | Go | 94 | 14 | 28 | 136 |
| [dir/funcdir_test.go](/dir/funcdir_test.go) | Go | 780 | 0 | 7 | 787 |
| [dir/vardir.go](/dir/vardir.go) | Go | 64 | 0 | 22 | 86 |
| [dir/vardir_test.go](/dir/vardir_test.go) | Go | 290 | 0 | 10 | 300 |
| [examples/balloon.lsh](/examples/balloon.lsh) | Lambdish | 141 | 2 | 10 | 153 |
| [examples/filter.lsh](/examples/filter.lsh) | Lambdish | 15 | 0 | 1 | 16 |
| [examples/length.lsh](/examples/length.lsh) | Lambdish | 6 | 0 | 0 | 6 |
| [examples/map.lsh](/examples/map.lsh) | Lambdish | 10 | 0 | 1 | 11 |
| [examples/maplist.lsh](/examples/maplist.lsh) | Lambdish | 26 | 1 | 2 | 29 |
| [examples/mergesort.lsh](/examples/mergesort.lsh) | Lambdish | 56 | 1 | 7 | 64 |
| [examples/oddmatrix.lsh](/examples/oddmatrix.lsh) | Lambdish | 143 | 2 | 17 | 162 |
| [examples/reduce.lsh](/examples/reduce.lsh) | Lambdish | 15 | 0 | 1 | 16 |
| [examples/reverse.lsh](/examples/reverse.lsh) | Lambdish | 10 | 0 | 1 | 11 |
| [examples/sortparity.lsh](/examples/sortparity.lsh) | Lambdish | 23 | 1 | 4 | 28 |
| [examples/test_algorithms.lsh](/examples/test_algorithms.lsh) | Lambdish | 99 | 12 | 11 | 122 |
| [examples/test_error.lsh](/examples/test_error.lsh) | Lambdish | 19 | 6 | 5 | 30 |
| [examples/test_lambdas.lsh](/examples/test_lambdas.lsh) | Lambdish | 41 | 5 | 5 | 51 |
| [examples/test_matrix.lsh](/examples/test_matrix.lsh) | Lambdish | 46 | 5 | 5 | 56 |
| [go.mod](/go.mod) | XML | 6 | 0 | 3 | 9 |
| [gocc/compilegocc.sh](/gocc/compilegocc.sh) | Shell Script | 1 | 1 | 1 | 3 |
| [gocc/errors/errors.go](/gocc/errors/errors.go) | Go | 48 | 1 | 8 | 57 |
| [gocc/lexer/acttab.go](/gocc/lexer/acttab.go) | Go | 235 | 1 | 8 | 244 |
| [gocc/lexer/lexer.go](/gocc/lexer/lexer.go) | Go | 96 | 65 | 17 | 178 |
| [gocc/lexer/transitiontable.go](/gocc/lexer/transitiontable.go) | Go | 628 | 61 | 4 | 693 |
| [gocc/parser/action.go](/gocc/parser/action.go) | Go | 41 | 1 | 10 | 52 |
| [gocc/parser/actiontable.go](/gocc/parser/actiontable.go) | Go | 6,139 | 1 | 4 | 6,144 |
| [gocc/parser/gototable.go](/gocc/parser/gototable.go) | Go | 3,640 | 1 | 5 | 3,646 |
| [gocc/parser/parser.go](/gocc/parser/parser.go) | Go | 180 | 4 | 33 | 217 |
| [gocc/parser/productionstable.go](/gocc/parser/productionstable.go) | Go | 390 | 2 | 5 | 397 |
| [gocc/token/token.go](/gocc/token/token.go) | Go | 133 | 15 | 23 | 171 |
| [gocc/util/litconv.go](/gocc/util/litconv.go) | Go | 85 | 12 | 12 | 109 |
| [gocc/util/rune.go](/gocc/util/rune.go) | Go | 35 | 1 | 4 | 40 |
| [grammar/grammar_test.go](/grammar/grammar_test.go) | Go | 44 | 0 | 14 | 58 |
| [grammar/lambdish.ebnf](/grammar/lambdish.ebnf) | EBNF | 83 | 0 | 20 | 103 |
| [grammar/tests/test1.lsh](/grammar/tests/test1.lsh) | Lambdish | 3 | 0 | 2 | 5 |
| [grammar/tests/test2.lsh](/grammar/tests/test2.lsh) | Lambdish | 13 | 0 | 2 | 15 |
| [grammar/tests/test3.lsh](/grammar/tests/test3.lsh) | Lambdish | 13 | 0 | 2 | 15 |
| [grammar/tests/test4.lsh](/grammar/tests/test4.lsh) | Lambdish | 12 | 0 | 1 | 13 |
| [grammar/tests/test5.lsh](/grammar/tests/test5.lsh) | Lambdish | 7 | 0 | 2 | 9 |
| [grammar/tests/test6.lsh](/grammar/tests/test6.lsh) | Lambdish | 11 | 0 | 3 | 14 |
| [ic/addressstack.go](/ic/addressstack.go) | Go | 54 | 9 | 17 | 80 |
| [ic/generateaddresses.go](/ic/generateaddresses.go) | Go | 104 | 7 | 25 | 136 |
| [ic/generateaddresses_test.go](/ic/generateaddresses_test.go) | Go | 39 | 2 | 10 | 51 |
| [ic/generatecode.go](/ic/generatecode.go) | Go | 493 | 41 | 135 | 669 |
| [ic/generator.go](/ic/generator.go) | Go | 117 | 20 | 35 | 172 |
| [ic/ic.go](/ic/ic.go) | Go | 35 | 12 | 11 | 58 |
| [ic/ic_test.go](/ic/ic_test.go) | Go | 58 | 0 | 17 | 75 |
| [ic/tests/test1.lsh](/ic/tests/test1.lsh) | Lambdish | 11 | 0 | 2 | 13 |
| [ic/tests/test2.lsh](/ic/tests/test2.lsh) | Lambdish | 15 | 0 | 2 | 17 |
| [ic/tests/test3.lsh](/ic/tests/test3.lsh) | Lambdish | 16 | 0 | 5 | 21 |
| [ic/tests/test4.lsh](/ic/tests/test4.lsh) | Lambdish | 11 | 0 | 1 | 12 |
| [ic/tests/test5.lsh](/ic/tests/test5.lsh) | Lambdish | 9 | 0 | 1 | 10 |
| [ic/tests/test6.lsh](/ic/tests/test6.lsh) | Lambdish | 9 | 0 | 1 | 10 |
| [ic/tests/test7_fake.lsh](/ic/tests/test7_fake.lsh) | Lambdish | 94 | 1 | 11 | 106 |
| [ic/util.go](/ic/util.go) | Go | 43 | 0 | 11 | 54 |
| [integration/integration_test.go](/integration/integration_test.go) | Go | 65 | 0 | 19 | 84 |
| [integration/test1.lsh](/integration/test1.lsh) | Lambdish | 9 | 0 | 1 | 10 |
| [integration/test2.lsh](/integration/test2.lsh) | Lambdish | 20 | 0 | 2 | 22 |
| [integration/test3.lsh](/integration/test3.lsh) | Lambdish | 23 | 0 | 3 | 26 |
| [mem/mem.go](/mem/mem.go) | Go | 290 | 46 | 42 | 378 |
| [mem/util.go](/mem/util.go) | Go | 5 | 0 | 2 | 7 |
| [quad/quad.go](/quad/quad.go) | Go | 229 | 7 | 22 | 258 |
| [sem/funccheck.go](/sem/funccheck.go) | Go | 104 | 5 | 25 | 134 |
| [sem/funccheck_test.go](/sem/funccheck_test.go) | Go | 36 | 0 | 10 | 46 |
| [sem/funcutil.go](/sem/funcutil.go) | Go | 15 | 0 | 5 | 20 |
| [sem/scopecheck.go](/sem/scopecheck.go) | Go | 88 | 11 | 18 | 117 |
| [sem/scopecheck_test.go](/sem/scopecheck_test.go) | Go | 39 | 2 | 11 | 52 |
| [sem/scopeutil.go](/sem/scopeutil.go) | Go | 19 | 4 | 6 | 29 |
| [sem/sem.go](/sem/sem.go) | Go | 19 | 25 | 7 | 51 |
| [sem/sem_test.go](/sem/sem_test.go) | Go | 53 | 2 | 16 | 71 |
| [sem/semanticcube.go](/sem/semanticcube.go) | Go | 240 | 7 | 34 | 281 |
| [sem/tests/test1.lsh](/sem/tests/test1.lsh) | Lambdish | 1 | 0 | 0 | 1 |
| [sem/tests/test2.lsh](/sem/tests/test2.lsh) | Lambdish | 7 | 0 | 1 | 8 |
| [sem/tests/test3.lsh](/sem/tests/test3.lsh) | Lambdish | 13 | 0 | 2 | 15 |
| [sem/tests/test4.lsh](/sem/tests/test4.lsh) | Lambdish | 11 | 0 | 1 | 12 |
| [sem/tests/test5.lsh](/sem/tests/test5.lsh) | Lambdish | 9 | 0 | 2 | 11 |
| [sem/tests/test6.lsh](/sem/tests/test6.lsh) | Lambdish | 11 | 0 | 3 | 14 |
| [sem/tests/test7_fake.lsh](/sem/tests/test7_fake.lsh) | Lambdish | 23 | 0 | 3 | 26 |
| [sem/typecheck.go](/sem/typecheck.go) | Go | 94 | 6 | 25 | 125 |
| [sem/typecheck_test.go](/sem/typecheck_test.go) | Go | 43 | 2 | 13 | 58 |
| [sem/typeutil.go](/sem/typeutil.go) | Go | 153 | 25 | 33 | 211 |
| [types/types.go](/types/types.go) | Go | 81 | 46 | 25 | 152 |
| [types/types_test.go](/types/types_test.go) | Go | 172 | 0 | 8 | 180 |
| [vm/ar/activationrecord.go](/vm/ar/activationrecord.go) | Go | 246 | 0 | 64 | 310 |
| [vm/ar/arstack.go](/vm/ar/arstack.go) | Go | 30 | 4 | 9 | 43 |
| [vm/list/list.go](/vm/list/list.go) | Go | 527 | 0 | 103 | 630 |
| [vm/list/liststack.go](/vm/list/liststack.go) | Go | 30 | 4 | 9 | 43 |
| [vm/mem.go](/vm/mem.go) | Go | 258 | 21 | 21 | 300 |
| [vm/operations.go](/vm/operations.go) | Go | 642 | 23 | 163 | 828 |
| [vm/tests/balloon.lsh](/vm/tests/balloon.lsh) | Lambdish | 141 | 1 | 9 | 151 |
| [vm/tests/failtest.lsh](/vm/tests/failtest.lsh) | Lambdish | 1 | 0 | 0 | 1 |
| [vm/tests/filter.lsh](/vm/tests/filter.lsh) | Lambdish | 15 | 0 | 1 | 16 |
| [vm/tests/lambdas.lsh](/vm/tests/lambdas.lsh) | Lambdish | 21 | 4 | 3 | 28 |
| [vm/tests/map.lsh](/vm/tests/map.lsh) | Lambdish | 10 | 0 | 1 | 11 |
| [vm/tests/mergesort.lsh](/vm/tests/mergesort.lsh) | Lambdish | 56 | 0 | 6 | 62 |
| [vm/tests/oddmatrix.lsh](/vm/tests/oddmatrix.lsh) | Lambdish | 94 | 1 | 11 | 106 |
| [vm/tests/reverse.lsh](/vm/tests/reverse.lsh) | Lambdish | 10 | 0 | 1 | 11 |
| [vm/tests/test1.lsh](/vm/tests/test1.lsh) | Lambdish | 10 | 0 | 2 | 12 |
| [vm/tests/test2.lsh](/vm/tests/test2.lsh) | Lambdish | 28 | 0 | 2 | 30 |
| [vm/tests/test3.lsh](/vm/tests/test3.lsh) | Lambdish | 20 | 0 | 6 | 26 |
| [vm/tests/test4.lsh](/vm/tests/test4.lsh) | Lambdish | 25 | 0 | 3 | 28 |
| [vm/tests/test5.lsh](/vm/tests/test5.lsh) | Lambdish | 13 | 0 | 3 | 16 |
| [vm/tests/test6.lsh](/vm/tests/test6.lsh) | Lambdish | 25 | 0 | 2 | 27 |
| [vm/tests/test7.lsh](/vm/tests/test7.lsh) | Lambdish | 21 | 0 | 3 | 24 |
| [vm/util.go](/vm/util.go) | Go | 184 | 2 | 49 | 235 |
| [vm/vm.go](/vm/vm.go) | Go | 291 | 7 | 42 | 340 |
| [vm/vm_test.go](/vm/vm_test.go) | Go | 34 | 1 | 11 | 46 |

[summary](results.md)