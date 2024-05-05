# auto-pdf-generator
自動でアプリを起動&自動で動かしながら スクリーンショットを撮ってPDFを生成します

https://github.com/ryomak/auto-pdf-generator/assets/21288308/f6316c14-067a-45ea-9247-4f681938d2c2

1. 指定したアプリを起動
2. スクショする画面の左上と右下の座標を指定
3. 自動でスクリーンショットを撮りながらスクロール
4. PDFを生成

※ 音声ガイド付きで進捗状況も教えてくれます

# 使い方
## インストール
```bash
git clone git@github.com:ryomak/screenshot-pdf-go.git
```
## CLI
### 実行
```bash
make build-cli
./bin/pdfgenerator-cli -o output.pdf -p 100 -s -a Chrome
```

## Desktop
### 実行
```bash
make install
make build-desktop
./bin/pdfgenerator-desktop
```

# NOTICE
- このツールはmacOSでのみ動作確認をしてます
- デフォルトではKindleアプリを開くようになっています。Amazonのカスタマーサポートに問い合わせ「個人利用であれば問題ない」ことを確認しております。
- 利用時に発生した損害について一切責任を負いません。ツールを使用する際は自己責任でお願いします。

![スクリーンショット 2024-05-03 23 04 00](https://github.com/ryomak/auto-pdf-generator/assets/21288308/b61c08e0-229c-4568-9d56-bf9f16ede6a8)

