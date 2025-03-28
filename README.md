# Dynamic MacroBoard (SendInput)

**Dynamic MacroBoard** is a **Windows-only** macro keyboard built in **Go** using **Ebiten** for the graphical interface and the **Windows API (SendInput)** for simulating global Unicode key presses.  
With this tool, you can configure keys that—when clicked—simulate **key events** even if the Ebiten window is **unfocused**! 🎯

---

## ✨ Features

✅ **Graphical Macro Keyboard** 🎨  
A simple UI built with [Ebiten](https://ebiten.org/) that displays clickable keys.

✅ **Global Key Emulation** 🌍  
Uses the **Windows API SendInput** function to simulate Unicode key events so that key presses are sent to the active window (for example, a browser).

✅ **Dynamic Key Configuration** 🔧  
Supports dynamically mapping **key labels to Unicode events**. You can extend this to load key mappings from a **configuration file**.

---

## 🛠️ Requirements

- **🖥️ Operating System:**  
Windows **64-bit** (this project uses Windows-specific API calls).

- **🐹 Go:**  
  Go **1.16+** is recommended.

- **📦 Dependencies:**
    - 🎮 [Ebiten](https://github.com/hajimehoshi/ebiten) (Game Library)
    - 📄 [YamlV3](https://github.com/go-yaml/yaml) (Config Parsing)

  No external packages are required for key simulation since we use `syscall` and `unsafe` to call the **Windows API** directly.

---


## 🛠️ How It Works

- **🎹 SendInput Integration:**  
  The project calls the **Windows API function `SendInput`** to simulate key events. The code defines custom `INPUT` and `KEYBDINPUT` structures with padding (to ensure correct size on **64-bit systems**) and determines the structure size based on the target architecture.

- **🖼️ Ebiten UI:**  
  **Ebiten** is used to create a window with clickable **buttons**. Each button represents a **macro key**, and clicking one triggers its associated handler.

- **⚡ Dynamic Key Handling:**  
  The handler (`EmulateKeyHandler`) loops through each **rune in the key’s label** and sends it using the `sendUnicodeChar` function.

