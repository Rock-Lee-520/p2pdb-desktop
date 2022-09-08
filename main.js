const { app, BrowserWindow,Menu } = require('electron')
const isDev = require('electron-is-dev')

// 我们这里new了两个 BrowserWindow，这是不行的，需要进一步封装
class AppWindow extends BrowserWindow {
  constructor(config) {
    const baseConfig = {
      width: 800,
      height: 600,
      webPreferences: {
        nodeIntegration: true
      }
    }
    // const finalConfig = Object.assign(baseConfig, config)
    const finalConfig = {...baseConfig, ...config}
    super(finalConfig)
    this.webContents.openDevTools()
  }
}


// 创建 menu
function createMenu() {
  let menuStructure = [
      {
          label: '配置',
          submenu: [
              {
                  label: '配置',
                  click() {
                      createConfigWindow()
                  }
              },
              {
                  label: '刷新', // 刷新页面
                  click() {
                      refreshWindows()
                  }
              },
              {
                  label: '打开调试窗口',
                  click(menuItem, targetWindow) {
                      targetWindow.openDevTools()
                  }
              },
              {
                  label: '关闭调试窗口',
                  click(menuItem, targetWindow) {
                      targetWindow.closeDevTools()
                  }
              },
          ]
      },
      // {
      //     label: '编辑',
      //     role: 'editMenu'
      // },
      {
          label: '文件夹',
          submenu: [
              {label: '打开 Rime 配置文件夹', click() {shell.openPath(getRimeConfigDir())}},
              {label: '打开 Rime 程序文件夹', click() {shell.openPath(getRimeExecDir())}},
              {
                  label: '打开工具配置文件夹', click() {
                      let configDir = path.join(os.homedir(), CONFIG_FILE_PATH)
                      shell.openPath(configDir)
                  }
              },
          ]
      },
      // {
      //     label: '码表处理工具',
      //     submenu: [
      //         {
      //             label: '码表处理工具',
      //             click() {
      //                 showToolWindow()
      //             }
      //         },
      //     ]
      // },
      {
          label: '关于',
          submenu: [
              {label: '最小化', role: 'minimize'},
              {label: '关于', role: 'about'},
              {type: 'separator'},
              {label: '退出', role: 'quit'},
          ]
      },
  ]
  let menu = Menu.buildFromTemplate(menuStructure)
  Menu.setApplicationMenu(menu)
}


app.on('ready', () => {
  const urlLocation = isDev ? 'http://127.0.0.1:5200' : 'are you ok'
  const mainWindow = new AppWindow({});
  mainWindow.loadURL(urlLocation);
  createMenu(); // 创建菜单
})
