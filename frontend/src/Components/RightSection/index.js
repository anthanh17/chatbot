import { useState } from "react";
import { CautionIcon, LightningChargeIcon, SunIcon } from "../../constants";
import Footer from "../Footer";
import ChatList from "./ChatList";
import axios from "axios";

const SUGGESTION = [
  {
    icon: <SunIcon />,
    title: "Examples",
    subTitle: [
      `"Explain quantum computing in simple terms" →`,
      `"Got any creative ideas for a 10 year old’s birthday?" →`,
      `"How do I make an HTTP request in Javascript?" →`,
    ],
    hover: true,
  },
  {
    icon: <LightningChargeIcon />,
    title: "Capabilities",
    subTitle: [
      `Remembers what user said earlier in the conversation`,
      `Allows user to provide follow-up corrections`,
      `Trained to decline inappropriate requests`,
    ],
    hover: false,
  },
  {
    icon: <CautionIcon />,
    title: "Limitations",
    subTitle: [
      `May occasionally generate incorrect information`,
      `May occasionally produce harmful instructions or biased content`,
      `Limited knowledge of world and events after 2021`,
    ],
    hover: false,
  },
];

const fetchDataPrompt = async (question) => {
  const apiUrl = "http://localhost:9000/chatbot/prompt";
  const postData = {
    question: question,
  };

  console.log("postData", postData);

  try {
    const response = await axios.post(apiUrl, postData, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.error("Annt RightSection", response.data);

    return response.data.answer;
  } catch (error) {
    console.error("API Error:", error);
    // Handle API errors here
    return null;
  }
};

const RightSection = () => {
  const [chatList, setChatList] = useState([]);

  const onSubmitText = async (chat) => {
    setChatList((chatList) => [
      ...chatList,
      { user: "Me", message: chat, messageId: Math.random() },
    ]);

    console.error("Annt RightSection", chat);
    const res = await fetchDataPrompt(chat);
    console.log(res);
    if (res)
      // check respone api success ở đây
      setChatList((chatList) => [
        ...chatList,
        { user: "GPT", message: res, messageId: Math.random() },
      ]);
  };

  return (
    <div className="flex h-full flex-1 flex-col md:pl-[260px]">
      <main className="relative h-full w-full transition-width flex flex-col overflow-hidden items-stretch flex-1">
        <div className="flex-1 overflow-hidden bg-lightBlack h-full">
          {!chatList.length ? (
            <div className="flex flex-col items-center text-sm h-full md:h-screen">
              <div className="text-gray-800 w-full md:max-w-2xl lg:max-w-3xl md:h-full md:flex md:flex-col px-6">
                <h1 className="text-4xl text-gray-100 font-semibold text-center mt-6 sm:mt-[20vh] ml-auto mr-auto mb-10 sm:mb-16">
                  Chatbot Ai
                </h1>
                <div className="md:flex items-start text-center gap-3.5">
                  {SUGGESTION.map((item, index) => (
                    <div
                      className="flex flex-col mb-8 md:mb-auto gap-3.5 flex-1"
                      key={index}>
                      <h2 className="flex gap-3 text-gray-100 items-center m-auto text-lg font-normal md:flex-col md:gap-2">
                        {item.icon}
                        {item.title}
                      </h2>
                      <ul className="flex flex-col gap-3.5 w-full sm:max-w-md m-auto">
                        {item.subTitle.map((subTitle, subTitleIndex) => (
                          <button
                            className={`w-full bg-gray-50 text-gray-100 bg-white/5 p-3 rounded-md ${
                              item.hover
                                ? "hover:bg-gray-200 dark:hover:bg-gray-900 cursor-pointer"
                                : "cursor-text"
                            }`}
                            key={subTitleIndex}>
                            {subTitle}
                          </button>
                        ))}
                      </ul>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          ) : (
            <ChatList chatList={chatList} />
          )}
        </div>
        <Footer onSubmitText={onSubmitText} />
      </main>
    </div>
  );
};

export default RightSection;
