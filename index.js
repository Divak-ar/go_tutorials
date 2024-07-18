import express, { urlencoded } from "express";

const app = express();

// db connection

const port = 3000;

// middle wares

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
// for text that if sent by user
app.use(express.text())

// routes

app.get("/", (req, res) => {
  res.status(200).send("Welcome to Go tutorials!!!");
});

app.get("/get", (req, res) => {
  res.status(200).json({ message: "Hello from the other side" });
});

app.post("/post", (req, res) => {
  let body = req.body;
  console.log(body)

  res.status(200).send(body)
});


// for from data -> in postman send in form-urlencode data
app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
  console.log(`Server Started at : localhost: http://localhost:${port}`);
});
