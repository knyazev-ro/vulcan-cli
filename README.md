# **gerard**

# 🧱 Gerard — Go CLI tool to rapidly scaffold clean, opinionated microservices. 

**Keep your architecture clean. Keep your team fast.**

**gerard** is a lightweight and developer-friendly CLI tool for quickly scaffolding API modules using the MVC pattern in Go. Powered by [Gorilla MUX](https://github.com/gorilla/mux) under the hood.

With **gerard**, you can initialize new modules, add controllers, and inject middleware in just one command.

---

## 🚀 Features

✅ Scaffold modular MVC architecture  
✅ Add controllers with pre-generated handler methods  
✅ Inject middleware with registration stubs  
✅ Works out-of-the-box — no dependencies beyond the standard Go toolchain  
✅ Designed for rapid web API prototyping

---

## 🛠️ Getting Started

### 📦 Step 1: Download or clone `gerard` into the root of your Go project

```bash
git clone https://github.com/knyazev-ro/gerard-cli.git
````

*Or just copy the `gerard/` folder into your Go project.*

---

### ⚙️ Step 2: Build the CLI

```bash
cd gerard
go build -o ../gerard.exe .
cd ..
```

You should now see `gerard.exe` in your root project directory.

---

### 🎉 Step 3: Use the CLI!

Run any of the following commands from your project root:

#### 🔧 Initialize a new module

```bash
gerard.exe create:init blog
```

Creates a new `blog` module with folders, routes, and sample boilerplate.

#### 🧩 Add middleware to a module

```bash
gerard.exe create:middleware auth blog
```

Creates a new `authMiddleware.go` in the `blog/middleware/` folder and registers it automatically.

#### 🧠 Add a controller to a module

```bash
gerard.exe create:controller post blog
```

Generates `postController.go` with empty handlers (e.g., `GetPosts`, `CreatePost`, etc.) and auto-wires it into the router.

---

## 📁 Example Structure After Running

```
your_project/
├── gerard.exe
├── blog/
|   src/
│   ├── controller/
│   │   └── postController.go
│   ├── middleware/
│   │   └── authMiddleware.go
|   |── routes/
│   |   └── routes.go
|   └──main.go
```

---

## 📌 Requirements

* Go 1.18+
* Windows (for `.exe`) — or just build without `-o` for other platforms

---

## 🗂️ What's Next?

In the future, we plan to support:

* Automatic model generation
* Swagger doc comments
* CLI help & autocomplete

---

## 📄 License

MIT

---

