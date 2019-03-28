## 常見的 HTTP Method：
  轉貼：[常見的HTTP METHOD](https://data-sci.info/2015/10/24/%E5%B8%B8%E8%A6%8B%E7%9A%84http-method%E7%9A%84%E4%B8%8D%E5%90%8C%E6%80%A7%E8%B3%AA%E5%88%86%E6%9E%90%EF%BC%9Agetpost%E5%92%8C%E5%85%B6%E4%BB%964%E7%A8%AEmethod%E7%9A%84%E5%B7%AE%E5%88%A5/)
  
  <br/>
    文中列舉出常見的 HTTP Method 有六種 HTTP Method分別是 **head,get,post,delete,put,patch**
    <br/>
    不同的Method就是對同一件事情做不同的操作。
    舉服務生點餐的例子，
    假設現在我們要點餐，我們必須先知道菜單是甚麼(get)，
    我們會向服務生點餐（post），
    我們想要取消剛才點的餐點（delete），
    我們想要重新點一次（put），
    我們想要加點甜點和飲料（patch）。

    head是取得get的http header而不取得內容，性質上我們可以當作跟get一樣），至於這幾種Method的行為我們可以統整一下：
 
    head：和get一樣，只是head只會取的HTTP header的資料。
    get：取得我們想要的資料。
    post：新增一項資料。（如果存在會新增一個新的）
    put：新增一項資料，如果存在就覆蓋過去。（還是只有一筆資料，跟服務生說「我」要吃牛排說10次還是只會來一份牛排）
    patch：附加新的資料在已經存在的資料後面。（資料必須已經存在，patch會擴充這項資料）
    delete：刪除資料。
    
| Methods   |      Safe?（是否安全)     | Idempotent?（是否可以重新整理） |
|----------|:-------------:|------:|
| GET |  Y | Y |
| POST |    N   |   N |
| PATCH | N |    N |
|PUT | N | Y|
|DELETE | N | Y
