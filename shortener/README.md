Написать тип, который реализует интерфейс:

```golang
type Shortener interface {
    Shorten(url string) string
    Resolve(url string) string
}
```

Метод Shorten - возвращает "короткую" ссылку (выбор алгоритма - за студентом), например otus.ru/some-long-link -> otus.ru/jhg34 и сохраняет соответствие короткой и исходной ссылок в памяти (не используя БД, а использовать, например, map).
При вызове метода Resolve - отдавает "длинную ссылку" или пустую строку, если ссылка не найдена.