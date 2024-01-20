print("digite o nome do programa:")
local nome = io.read()

print("digite o path do executavel:")
local idade = io.read()

print("digite o path do icone:")
local icone = io.read()

print("digite o nome do arquivo:")
local namefile = io.read()


function CreateDesktopEntry(name, exec, icon, namefile)
    local content = string.format(
        "[Desktop Entry]\n" ..
        "Name=%s\n" ..
        "Exec=%s\n" ..
        "Icon=%s\n" ..
        "Type=Application\n" ..
        "Categories=Application;\n",
        name, exec, icon
    )

    local file = io.open(namefile .. ".desktop", "w")

    if file then
        file:write(content)
        file:close()
    end

end

createDesktopEntry(nome, idade, icone, namefile)