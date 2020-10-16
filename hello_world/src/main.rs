use sincere::App;

fn main() {
    let mut app = App::new();

    app.get("/", |context| {
        context.response.from_json("{\"name\": \"Hello World\"}").unwrap();
    });

    app.run("127.0.0.1:8000").unwrap()
}
