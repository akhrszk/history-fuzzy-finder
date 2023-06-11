# history-fuzzy-finder

.zsh_historyからコマンド履歴を検索するfuzzy finderをGoで作ってみた。

![output](https://github.com/akhrszk/history-fuzzy-finder/assets/28677705/1f8f7711-ba4b-48e8-b412-01cbde85856a)

## 試してみる

```
go run main.go | xargs echo
```

## Dependencies

- zsh（.zsh_historyからコマンド履歴を検索するのでzshが使われていることが前提）
- [rivo/tview](https://github.com/rivo/tview)（UI部分はtviewで実装）
