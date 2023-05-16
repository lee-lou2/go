const token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJsZWVAbG91Mi5rciIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlfSwiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS9hdXRoIjp7InVzZXJfaWQiOiJ1c2VyLWw4ZWNWbjNzS1lPVGtMdHU0TFRqVU5aRCJ9LCJpc3MiOiJodHRwczovL2F1dGgwLm9wZW5haS5jb20vIiwic3ViIjoiZ29vZ2xlLW9hdXRoMnwxMDQwMjgwNzIxMjE2ODYzNDMzNzYiLCJhdWQiOlsiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS92MSIsImh0dHBzOi8vb3BlbmFpLm9wZW5haS5hdXRoMGFwcC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjgzMjM3MDk2LCJleHAiOjE2ODQ0NDY2OTYsImF6cCI6IlRkSkljYmUxNldvVEh0Tjk1bnl5d2g1RTR5T282SXRHIiwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBtb2RlbC5yZWFkIG1vZGVsLnJlcXVlc3Qgb3JnYW5pemF0aW9uLnJlYWQgb2ZmbGluZV9hY2Nlc3MifQ.FZIlMqHnfegzXb-9HNJ__6TNzUOQZBo1UxuEBVCEZ8zrhaxm8h0PgazamoT2W6xY84C27QUtOEDRh1Ud71Vz6cp8j3OOT4TJ9qMNHiDTsc6cxVIjNBSZpY7QhCfIwarhIimDz24PtPM9ViCWzoJGwwKYcYj9uM-LfJZDZr7vUdbo1JF-wCXiM2bwpUfzraaHIhfI2ZiRakNAdgKrDPVjwgQ7IRFEJkYlpTWw-WaEkFfbBYmIYILt8AgeBk1Em4J6z5coWc6C3M1bMLT7aK_0FmgKeV2D3DK4skiJdeRIiMVpMcBNcyWMTRYAa9HCWCwiN0UuZ_xLHmhuPBgQakjZQw";
const server = "https://api.lou2.kr";


function updateOutput(objId, message) {
    fetch(server + "/v1/ai/dataset/" + objId, {
        "headers": {"content-type": "application/json"},
        "method": "PATCH",
        "body": JSON.stringify({
            "output": message
        })
    })
}

function sendSlack(category, message) {
    if (category.indexOf("slack") == -1) {
        console.log("슬랙으로 데이터 전송하지 않음");
        return
    }
    console.log("슬랙으로 데이터 전송");
    const url = server + "/v1/notify/slack";
    // post fetch 전송
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({data: {
                channel: "chatgpt",
                message: message
            }})
    }).then(res => {
        if (res.status !== 200) {
            throw new Error('Error : ' + res.status);
        }
    })
}

async function training() {
    let resp = await fetch(server + "/v1/ai/dataset/instruction");
    if (resp.status != 200) {
        console.log("조회된 데이터가 없습니다.")
        return
    }
    let obj = await resp.json();
    let body = "{\"action\":\"next\",\"messages\":[{\"id\":\"aaa26343-7a66-4d0d-9085-794b01b529d3\",\"author\":{\"role\":\"user\"},\"content\":{\"content_type\":\"text\",\"parts\":[\"__message__\"]}}],\"parent_message_id\":\"aaa14326-ba0a-4acf-8d72-1fb1cab11c25\",\"model\":\"text-davinci-002-render-sha\",\"timezone_offset_min\":-540,\"history_and_training_disabled\":false}";
    const headers = {
        "accept": "text/event-stream",
        "accept-language": "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7",
        "authorization": "Bearer " + token,
        "content-type": "application/json",
        "sec-ch-ua": "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": "\"macOS\"",
        "sec-fetch-dest": "empty",
        "sec-fetch-mode": "cors",
        "sec-fetch-site": "same-origin"
    };
    // __message__ 라는 문자열을 받은 메시지로 대체합니다.
    body = body.replace('__message__', obj.data.instruction);
    fetch("https://chat.openai.com/backend-api/conversation", {
        "headers": headers,
        "referrer": "https://chat.openai.com/?model=text-davinci-002-render-sha",
        "referrerPolicy": "same-origin",
        "body": body,
        "method": "POST",
        "mode": "cors",
        "credentials": "include"
    }).then(res => {
        if (res.status !== 200) {
            throw new Error('Error : ' + res.status);
        }
        // 스트리밍되는 데이터를 받아서 처리합니다.
        const reader = res.body.getReader();
        // 불러온 데이터를 리스트에 저장합니다.
        const stream = new ReadableStream({
            start(controller) {
                function push() {
                    // reader.read()는 Promise를 반환합니다.
                    return reader.read().then(({done, value}) => {
                        // 더 이상 받을 데이터가 없으면 종료합니다.
                        if (done) {
                            controller.close();
                            return;
                        }
                        // 받은 데이터를 컨트롤러에 전달합니다.
                        controller.enqueue(value);
                        return push();
                    });
                }
                return push();
            }
        });
        // 컨트롤러에서 스트림을 읽습니다.
        return new Response(stream, {headers: {"Content-Type": "text/html"}}).text();
    }).then(result => {
        // 위 텍스트를 파싱해서 필요한 데이터만 추출합니다. (JSON 형식으로 파싱)
        const json = JSON.parse(result.split('\n\n')[result.split('\n\n').length - 3].replace('data: ', ''));
        // 필요한 데이터를 추출합니다.
        const text = json.message.content.parts[0];
        // 소켓으로 전송
        console.log('Result : ' + text);
        updateOutput(obj.data.ID, text);
        sendSlack(obj.data.category, text)
        return json.conversation_id;
    }).then(message_id => {
        fetch("https://chat.openai.com/backend-api/conversation/" + message_id, {
            "headers": headers,
            "referrer": "https://chat.openai.com/c/" + message_id,
            "referrerPolicy": "same-origin",
            "body": "{\"is_visible\":false}",
            "method": "PATCH",
            "mode": "cors",
            "credentials": "include"
        }).then(res => {console.log("removed message");});
    }).catch(err => {
        console.error(err);
    });
}

setInterval(() => {
    training();
}, 5000);