const { app, BrowserWindow } = require('electron')
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

app.on('ready', () => {
  const urlLocation = isDev ? 'http://localhost:5000/' : 'are you ok'
  const mainWindow = new AppWindow({})
  mainWindow.loadURL(urlLocation)
})
