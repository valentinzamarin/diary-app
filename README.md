### Diary app 

Personal project. The initial version was built on pure Electron. After some time decided to move away from it but keep the frontend on the HTML/CSS/JS stack, as writing something aesthetically pleasing in Fyne is quite a challenge.

will be getting new features sometimes

Stack : 
- Golang
- SQLite / GORM
- Wails 
- Vue3 ( composition API )

dev:
```
wails dev -tags webkit2_41
```

build 
```
wails build -tags webkit2_41
```

---

tried to keep the frontend as simple as possible - anything that could be done via the backend was done via the backend ( wails events as replacement for vue features ). I avoided pulling in state management libraries, like Pinia, without a real need. The main focus was on the backend - the frontend is just there to look nice and for my own personal satisfaction.

