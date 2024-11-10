#### On a quest to catch 'em all: *node_modules edition*

##### :calling: Reach me at **[email](mailto:johannes@stenmark.in)** ***/*** **[LinkedIn](https://www.linkedin.com/in/johannes-stenmark)**.  :feet: Check out the [ChatCopyCat](https://github.com/jstenmark/ChatCopyCat) extension for VSCode.

---
#### :cookie: Fortune cookie of the day
```smalltalk
╭──────────────────────────────────────────────────────────────────────────────╮
│ There has also been some work to allow the interesting use of macro names.   │
│ For example, if you wanted all of your "creat()" calls to include read       │
│ permissions for everyone, you could say                                      │
│                                                                              │
│     #define creat(file, mode)    creat(file, mode | 0444)                    │
│                                                                              │
│     I would recommend against this kind of thing in general, since it        │
│ hides the changed semantics of "creat()" in a macro, potentially far away    │
│ from its uses.                                                               │
│     To allow this use of macros, the preprocessor uses a process that        │
│ is worth describing, if for no other reason than that we get to use one of   │
│ the more amusing terms introduced into the C lexicon.  While a macro is      │
│ being expanded, it is temporarily undefined, and any recurrence of the macro │
│ name is "painted blue" -- I kid you not, this is the official terminology    │
│ -- so that in future scans of the text the macro will not be expanded        │
│ recursively.  (I do not know why the color blue was chosen; I'm sure it      │
│ was the result of a long debate, spread over several meetings.)              │
│         -- From Ken Arnold's "C Advisor" column in Unix Review               │
│                                                                              │
╰──────────────────────────────────────────────────────────────────────────────╯
```
