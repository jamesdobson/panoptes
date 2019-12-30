import Foundation
import ScreenSaver
import WebKit


// MARK: - SaverView
final class SaverView: ScreenSaverView {
  
  private var prefs: UserDefaults? = nil
  
  private var urlIndex: Int = 0
  
  private var urlArray: [String] = []
    
  // MARK: Outlets
  private var webView: WKWebView

  // MARK: Initialization
  override init?(frame: NSRect, isPreview: Bool) {
    let webConfiguration = WKWebViewConfiguration()
    let bounds = NSRect(x: 0, y: 0, width: frame.width, height: frame.height)
    
    webView = WKWebView(frame: bounds, configuration: webConfiguration)

    super.init(frame: frame, isPreview: isPreview)

    configure()
}

  required init?(coder decoder: NSCoder) {
    let webConfiguration = WKWebViewConfiguration()
    
    webView = WKWebView(frame: .zero, configuration: webConfiguration)
    
    super.init(coder: decoder)

    configure()
  }
}

// MARK: - Lifecycle
extension SaverView {
  override func animateOneFrame() {
    super.animateOneFrame()
    
    showURL(url: urlArray[urlIndex % urlArray.count])
    urlIndex += 1
  }
  
  func showURL(url: String) {
    let req = URLRequest(url: URL(string:url)!)
    webView.load(req)
  }
}

// MARK: - Configuration
private extension SaverView {

  func configure() {
    prefs = ScreenSaverDefaults(forModuleWithName: "com.softwarepunk.Panoptes")!
    
    if let urls = prefs?.stringArray(forKey: "urls") {
      urlArray = urls
    } else {
      urlArray = ["https://www.apple.com"]
      prefs!.setValue(urlArray, forKey:  "urls")
      prefs!.synchronize()
    }
    
    let interval = prefs!.double(forKey: "intervalSecs")
    
    if interval > 0 {
      animationTimeInterval = interval
    } else {
      animationTimeInterval = 60.0
      prefs!.set(animationTimeInterval, forKey: "intervalSecs")
      prefs!.synchronize()
    }

    addSubviews()
    defineConstraints()
    setupSubviews()
  }
  
  func addSubviews() {
    addSubview(webView)
  }

  func defineConstraints() {
  }

  func setupSubviews() {
  }
}




