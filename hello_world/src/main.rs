use sincere::App;

fn main() {
    let mut app = App::new();

    app.get("/hello", |context| {
        context.response.from_text("Hello World").unwrap();
    });

    app.run("127.0.0.1:8000").unwrap()
}
