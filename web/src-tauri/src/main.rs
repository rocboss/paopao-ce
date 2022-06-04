#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

#[cfg(target_os = "macos")]
use tauri::{api::shell, CustomMenuItem, Manager, Menu, MenuEntry, MenuItem, Submenu};

fn main() {
    let _ctx = tauri::generate_context!();

    #[cfg(target_os = "macos")]
    let app = tauri::Builder::default()
        .menu(Menu::with_items([
            MenuEntry::Submenu(Submenu::new(
                &_ctx.package_info().name,
                Menu::with_items([
                    MenuItem::Separator.into(),
                    MenuItem::Services.into(),
                    MenuItem::Separator.into(),
                    MenuItem::Hide.into(),
                    MenuItem::HideOthers.into(),
                    MenuItem::ShowAll.into(),
                    MenuItem::Separator.into(),
                    MenuItem::Quit.into(),
                ]),
            )),
            MenuEntry::Submenu(Submenu::new(
                "Window",
                Menu::with_items([MenuItem::Minimize.into(), MenuItem::Zoom.into()]),
            )),
            // You should always have a Help menu on macOS because it will automatically
            // show a menu search field
            MenuEntry::Submenu(Submenu::new(
                "Help",
                Menu::with_items([CustomMenuItem::new("Learn More", "Learn More").into()]),
            )),
        ]))
        .on_menu_event(|event| {
            let event_name = event.menu_item_id();
            event.window().emit("menu", event_name).unwrap();
            match event_name {
                "Learn More" => {
                    let link = "https://github.com/rocboss/paopao-ce".to_string();
                    shell::open(&event.window().shell_scope(), link, None).unwrap();
                }
                _ => {}
            }
        });

    #[cfg(not(target_os = "macos"))]
    let app = tauri::Builder::default();

    app.run(tauri::generate_context!())
        .expect("error while running tauri application");
}
