import './App.css'
import {useState} from "react";
import {Execute} from "../wailsjs/go/main/App";

function App() {

    const [output, setOutput] = useState<string>('example.pdf')
    const [page, setPage] = useState<number>(1)
    const [nextAction, setNextAction] = useState<string>('down')
    const [split, setSplit] = useState<boolean>(false)
    const [appName, setAppName] = useState<string>('Kindle')
    const [slideDuration, setSlideDuration] = useState<number>(300)

    const [result, setResult] = useState<string>('')
    const changeEvent = (fn:any) => (e: React.ChangeEvent<HTMLInputElement>) => {
        return fn(e.target.value);
    }


    const handleExecute = async () => {
        try {
            const response = await Execute(output, page, nextAction, split, appName,slideDuration,"")
            setResult(response)
        } catch (e) {
            setResult('エラーが発生しました')
        }
    }

    return (
        <div className="min-h-screen bg-white grid grid-cols-1 place-items-center justify-items-center mx-auto py-8">
            <div className="text-blue-900 font-mono text-center mb-6">
                <h1 className="content-center text-2xl font-bold ">PDF Generator</h1>
                <p>起動したアプリを自動で動かしながら<br/>スクショを撮ってPDFに変換します</p>
            </div>
            <div className="w-fit max-w-md">
                {
                    !!result &&
                    <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-6" role="alert">
                        <strong className="font-bold">{result}</strong>
                    </div>
                }


                <form>
                    <div className="mb-6">
                        <label htmlFor="output" className="block mb-2 text-sm font-medium text-gray-900">出力先ファイル</label>
                        <input type="string" id="output"
                               className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                               placeholder="example.pdf" required value={output} onChange={changeEvent(setOutput)}/>
                    </div>
                    <div className="mb-6">
                        <label htmlFor="page" className="block mb-2 text-sm font-medium text-gray-900">最終ページ</label>
                        <input type="number" id="output"
                               className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                               placeholder="example.pdf" required value={page} onChange={(e)=>setPage(Number(e.target.value))}/>
                    </div>
                    <div className="mb-6">
                        <label htmlFor="action" className="block mb-2 text-sm font-medium text-gray-900">次へ進む時のアクション</label>
                        <select id="action"
                                onChange={(e)=>(setNextAction(e.target.value))}
                                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                            <option value="left">左</option>
                            <option value="right">右</option>
                            <option value="up">上</option>
                            <option value="down">下</option>
                        </select>
                    </div>

                    <div className="mb-6">
                        <label htmlFor="app" className="block mb-2 text-sm font-medium text-gray-900">起動アプリ</label>
                        <input type="string" id="output"
                               className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                               placeholder="example.pdf" required value={appName} onChange={changeEvent(setAppName)}/>
                    </div>
                    <div className="mb-6">
                        <label htmlFor="page" className="block mb-2 text-sm font-medium text-gray-900">移動時間(ms)</label>
                        <input type="number" id="output"
                               className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                               placeholder="example.pdf" required value={slideDuration} onChange={(e)=>setSlideDuration(Number(e.target.value))}/>
                    </div>
                    <div className="mb-6">
                        <label htmlFor="check" className="block mb-2 text-sm font-medium text-gray-900">画面を横に2分割するか</label>
                        <input type="checkbox" id="split"
                               className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                               placeholder="hogehoge"  onChange={(e)=> setSplit(e.target.checked)}/>
                    </div>
                    <button type="button"
                            onClick={handleExecute}
                            className="text-white bg-blue-900 hover:bg-blue-900 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-900 dark:hover:bg-blue-900 dark:focus:ring-blue-900">実行
                    </button>
                </form>
            </div>
        </div>
    )
}

export default App
