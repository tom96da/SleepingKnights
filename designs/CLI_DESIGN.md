# SleepingKnights CLI コマンド設計

**ステータス:** 提案段階（フェーズ別実装予定）
**最終更新:** 2026-07-14
**ターゲット:** 将来のコマンド体系設計とロードマップ

---

## 概要

SleepingKnights CLI は段階的に拡張される設計となっています。
本ドキュメントは **Phase 2以降の全コマンド体系** を定義しています。

---

## 1. Phase 2: 実行系コマンド

### 1.1 REPL（対話モード）

```bash
slk
```

コマンドラインで式や文を入力しながら対話的に実行します。

**参照:** [LANGUAGE_SPEC.md §11.1](LANGUAGE_SPEC.md#111-repl対話モード)

### 1.2 スクリプト実行

```bash
slk program.slk
slk ./src/main.slk
```

`.slk` 拡張子のスクリプトファイルをインタプリタモードで実行します。

### 1.3 コンパイルして実行

```bash
slk run program.slk
slk run ./src/main.slk
```

`.slk` 拡張子のスクリプトファイルをコンパイルしてから実行します。

**注:** コンパイルして実行は Phase 6 で提供予定です。

**参照:** [LANGUAGE_SPEC.md §11.3](LANGUAGE_SPEC.md#113-コンパイルして実行)

---

## 2. Phase 3: 検査・フォーマット系コマンド

### 2.1 構文チェック

```bash
slk check program.slk
slk check ./src                  # ディレクトリ内全ファイル
```

構文エラー + 型チェックエラーを報告（エラーのみ）

**出力例:**
```
src/program.slk:5:10: error: undefined variable 'x'
src/program.slk:7:3: error: type mismatch: int expected, string given
```

**オプション:**
- `--strict`: 警告も表示
- `--verbose`: 詳細情報表示

```bash
slk check --strict program.slk
```

### 2.2 フォーマッタ

```bash
slk fmt program.slk              # 整形差分を確認（変更しない）
slk fmt --fix program.slk        # ファイルを整形して上書き
slk fmt --fix ./src              # ディレクトリ内全ファイルを整形して上書き
```

SleepingKnights コード規約に従ったフォーマッティング

**フォーマット対象:**
- インデント（デフォルト: 2スペース）
- 演算子前後の空白
- 括弧内の空白
- 行末空白削除

**オプション:**
- `--fix`: 整形結果をファイルに適用
- `--indent=N`: インデント幅（デフォルト: 2）
- `--check`: 確認のみ（変更しない）

```bash
slk fmt --check program.slk      # 整形が必要か確認
slk fmt --fix --indent=4 program.slk   # 4スペースインデントで適用
```

---

## 3. Phase 3: 型チェック・リンター

### 3.1 リンター

```bash
slk lint program.slk
slk lint ./src --level=warn      # レベル指定
```

コード品質とベストプラクティス違反を報告

**検査項目（Phase 3で実装予定）:**
- 条件文での誤代入 (`if (x = 5)` → `if (x == 5)`)
- 未使用変数の警告
- 型推論矛盾
- 値を使わない関数呼び出し

**オプション:**
- `--level=error|warn|info`: 最小レポートレベル
- `--fix`: 自動修正を試みる（対応する警告のみ）
- `--config=.slklint`: 設定ファイル

```bash
slk lint --level=error program.slk      # エラーのみ表示
slk lint --fix program.slk              # 修正可能な警告を自動修正
```

---

## 4. Phase 6: コンパイル系コマンド

### 4.1 ネイティブコンパイル

```bash
slk build program.slk                    # program (or program.exe)
slk build -o myapp program.slk           # 出力ファイル名指定
slk build --mode=debug program.slk       # デバッグ情報付き
slk build --mode=release program.slk     # 最適化有効
```

LLVM を使用したネイティブバイナリへのコンパイル

**オプション:**
- `-o, --output FILE`: 出力ファイル名
- `--mode=debug|release`: ビルドモード（デフォルト: release）
- `--target=TRIPLE`: クロスコンパイル対象（例: `x86_64-unknown-linux-gnu`）
- `--verbose`: 詳細出力

```bash
slk build -o bin/app --mode=debug program.slk
```

### 4.2 ビルドチェーン

```bash
slk check program.slk               # 構文・型チェック
slk fmt --fix program.slk           # フォーマット
slk lint program.slk                # リンター
slk build -o myapp program.slk      # ビルド
```

---

## 5. Phase 7: ブートストラップ開始

ブートストラップ開始フェーズです（コンパイラの一部を自作言語で移植）。

**注:** この段階では専用CLIコマンドは未定です。

## 6. Phase 8: セルフホスティング達成

セルフホスティング達成フェーズです。

**注:** この段階では専用CLIコマンドは未定です。

---

## 7. 共通オプション

すべてのコマンドで利用可能：

| オプション | 説明 |
|-----------|------|
| `--help, -h` | ヘルプ表示 |
| `--version, -v` | バージョン表示 |
| `--verbose` | 詳細出力 |
| `--quiet, -q` | 静寂モード（エラーのみ） |

```bash
slk --version
slk run --help
slk fmt --verbose program.slk
```

---

## 8. 設定ファイル（将来拡張）

### 8.1 slk.toml

プロジェクトレベルの設定

```toml
[project]
name = "my-awesome-app"
version = "0.1.0"
description = "A simple SleepingKnights project"
authors = ["Your Name <you@example.com>"]

[build]
src = "src/main.slk"
output = "build/app"
mode = "release"

[format]
indent = 2
line_length = 100

[lint]
level = "warning"

[test]
pattern = "tests/**/*.test.slk"
```

### 8.2 .slklint

リンター設定（独立形式）

```toml
[warnings]
assignment_in_condition = true      # if (x = 5) 警告
unused_variables = true
unused_functions = false

[rules]
max_line_length = 100
```

---

## 9. 実装ロードマップ

| フェーズ | コマンド | 優先度 | 複雑度 |
|---------|---------|--------|--------|
| Phase 2 | `slk`, `slk file.slk` | 🔴 必須 | 低 |
| Phase 3 | `slk check` | 🟠 高 | 中 |
| Phase 3 | `slk fmt` | 🟠 高 | 中 |
| Phase 3 | `slk lint` | 🟡 中 | 高 |
| Phase 6 | `slk build` | 🔴 必須 | 高 |
| Phase 7 | ブートストラップ開始（専用CLI未定） | 🟢 低 | 中 |
| Phase 8 | セルフホスティング達成（専用CLI未定） | 🟢 低 | 中 |

---

## 10. 実装戦略

### 10.1 CLI フレームワーク

**推奨:** Go `cobra` ライブラリ

```go
// cmd/slk/cmd/root.go
var rootCmd = &cobra.Command{
  Use: "slk",
  Short: "SleepingKnights Programming Language",
}

// cmd/slk/cmd/run.go
var runCmd = &cobra.Command{
  Use: "run [file]",
  Run: func(cmd *cobra.Command, args []string) { ... },
}

// cmd/slk/cmd/check.go
var checkCmd = &cobra.Command{
  Use: "check [file|directory]",
  Run: func(cmd *cobra.Command, args []string) { ... },
}
```

**メリット:**
- 大規模なコマンド体系でも管理しやすい
- ヘルプ自動生成
- フラグ管理が簡潔
- コミュニティサポート豊富

**デメリット:**
- 依存関係が増える
- 学習コスト

### 10.2 段階的実装

1. **Phase 2**: `slk` / `slk file.slk` のシンプル実装（cobra 不要）
2. **Phase 3**: cobra 導入、`check`, `fmt`, `lint` 追加
3. **Phase 6**: `build`, `run`（コンパイルして実行）追加
4. **Phase 7**: ブートストラップ支援機能を追加
5. **Phase 8**: セルフホスティング検証機能を追加

---

## 11. エラーメッセージ例

### 構文エラー
```
src/program.slk:3:8: error: unexpected token ')'
  | let x = )
  |         ^
```

### 型エラー
```
src/program.slk:5:10: error: type mismatch
  expected: int
  actual:   string
  | let result = count + "5"
  |             ^^^^^^^^^^^^
```

### リンター警告
```
src/program.slk:7:5: warning: assignment in condition (linter)
  | if (x = 5) {
  |     ^^^^^^
  hint: use '==' for comparison, or 'if (let x = getValue())' for new variable binding
```

---

## 12. 今後の検討事項

- **デバッガ**: `slk debug` コマンド（Phase 8以降？）
- **パッケージ管理**: `slk add`, `slk remove` など（Phase 8以降？）
- **REPL 拡張**: `.help`, `.load`, `.save` メタコマンド（Phase 2以降）
- **カラー出力**: エラーメッセージのハイライト
- **並列処理**: ディレクトリ全体の並列チェック

---
