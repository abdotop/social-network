INSERT INTO users (id, email, password, first_name, last_name, date_of_birth, avatar_image, nickname, about_me, is_public) VALUES 
('a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'tester1@gmail.com', '$2a$10$H7v5cgz/oYCbJqskAg9bGukU6fAPAAi672Ki3W3u8OuNhhgESIL1e', 'Tester', 'Tested', datetime('2024-03-13'), 'uploads/default-avatar.png', 'testing1', 'not sure but think its the about me', 1),
('c035df0d-8699-4880-a79e-1291915f70a9', 'tester2@gmail.com', '$2a$10$H7v5cgz/oYCbJqskAg9bGukU6fAPAAi672Ki3W3u8OuNhhgESIL1e', 'Tester', 'Tested', datetime('2024-03-13'), 'uploads/default-avatar.png', 'testing2', 'not sure but think its the about me', 0),
('498e640d-78d2-4171-b060-369d75c380ed', 'tester3@gmail.com', '$2a$10$H7v5cgz/oYCbJqskAg9bGukU6fAPAAi672Ki3W3u8OuNhhgESIL1e', 'Tester', 'Tested', datetime('2024-03-13'), 'uploads/default-avatar.png', 'testing3', 'not sure but think its the about me', 1),
('36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'tester4@gmail.com', '$2a$10$H7v5cgz/oYCbJqskAg9bGukU6fAPAAi672Ki3W3u8OuNhhgESIL1e', 'Tester', 'Tested', datetime('2024-03-13'), 'uploads/default-avatar.png', 'testing4', 'not sure but think its the about me', 0),
('4f6f09ab-a290-4a38-9f70-9f61c2fd6e75', 'user@gmail.com', '$2a$10$H7v5cgz/oYCbJqskAg9bGukU6fAPAAi672Ki3W3u8OuNhhgESIL1e', 'User FirstName', 'UserName', datetime('2024-03-13'), 'uploads/3248494e-6de7-4de1-895a-1e8e59a03e00.jpeg', 'userNickName', 'Follow me... go on a tour and AMA', 1);
-- password: 12345678
-- PSEUDOS
-- '4f6f09ab-a290-4a38-9f70-9f61c2fd6e75', 
-- '94493d5c-08f0-4dd6-9d1f-e07e17316294', 
-- 'abb8c81e-c771-4d10-9d36-2c20e1faded2', 
-- '1cca8bec-260e-4241-9306-9c58e161bd4f', 
-- '2d00d01a-1882-41a1-97c0-58492472ea87', 

INSERT INTO followers (id, follower_id, followee_id, status) VALUES 
('16e526b0-022a-4dc1-8809-5fac78b5c8ca', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'requested'), 
('945206d2-1f56-445a-a064-11f4f8921a11', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '498e640d-78d2-4171-b060-369d75c380ed', 'accepted'),
('36acbd44-8c34-4357-a945-52695ab31f48', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'declined'), 
('bc4c42df-839c-44fa-8aae-b227ab64ed7f', '498e640d-78d2-4171-b060-369d75c380ed', 'c035df0d-8699-4880-a79e-1291915f70a9', 'requested'), 
('3b8dfc8c-4982-430f-abb4-7e95c809ad80', '498e640d-78d2-4171-b060-369d75c380ed', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'accepted');


INSERT INTO posts (id, user_id, title, content, image_url, privacy) VALUES
('3e8958e6-b6ac-4eea-ab67-2fe811211978', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'posting1', 'This is a post to populate the database', 'uploads/The+Real+Pros+&+Cons+of+Having+a+Nice+Car.png', 'public'), 
('8d91cc06-6847-484e-90d7-d19784a84a4a', 'c035df0d-8699-4880-a79e-1291915f70a9', 'post 02', 'Here another one to just populate again the database', 'uploads/b5a53ab198e4fc283d4666e8964ec772.jpg', 'private'), 
('9b77a08a-ff18-4022-97c7-4dc8f1a53230', '498e640d-78d2-4171-b060-369d75c380ed', 'posting more', 'At the end of the program, not only did I manage to catch up on my math skills, but I also discovered a new passion for this subject. I feel much more confident in my abilities, and my academic performance has improved significantly. I am grateful to all the teachers and organizers of this program for changing my perspective on mathematics and giving me the tools to succeed. I highly recommend this program to anyone who struggles with math or is looking to deepen their understanding of the subject.', '', 'almost private'),
('aec2952f-fa1c-4065-9e23-f660819ac4ff', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Apple Could Launch Its Most Important Product Since the iPhone This Summer, According to a Pair of Wall Street Analysts', 'Apple (NASDAQ: AAPL) could use a new blockbuster product, and it could come as soon as this summer.
The iPhone maker saw sales rise just 2% last quarter after a year in which revenue fell 2.8%. China remains a big question mark for Apple following a government ban of iPhones and a report indicating a massive decline in sales in the country to start the year. Meanwhile, the varying sources of its high-margin services revenue remain the focus of regulators around the world.
But at least two Wall Street analysts think Apple will release the next product to set off a spurt of growth at the massive company this summer. Melius Research analysts Ben Reitzes and Nick Monroe think investors should focus on what Apple will announce this summer at its annual Worldwide Developers Conference (WWDC).', 'uploads/05cb0f8594f61c5ad105f46f0d55dde7.webp', 'public'),
('9261e424-efd1-404c-96c8-76e9117d0606', '498e640d-78d2-4171-b060-369d75c380ed', 'India’s Quest for Building a Formidable Military-Industrial Complex', 'Rajnath Singh, the Defence Minister of India and the chief guest at the inaugural session of Firstpost Defence Summit held in New Delhi in Feb 2024, said India had to step out of its comfort zone to join the list of world’s top 25 arms exporters and Prime Minister Narendra Modi was focusing on long-term gains. India has exported Dornier-228s, 155 mm Advanced Towed Artillery Guns, Brahmos Missiles and Akash Missile System, while six countries are in talks to buy HAL’s indigenously manufactured light combat aircraft Tejas. “The aim is to manufacture high-end systems like aero-engines and gas turbines in India in the next five years,” he added. Rajnath Singh further said that the Modi Government is the first to halt the import of weapons to promote self-reliance. “We have made sure that our army uses indigenous resources and we even took a step forward to export these arms and equipment." 
According to Business Today, India’s arms exports touched Rs 16,000 crore in the 2022-23 financial year. Over 100 Indian companies are currently exporting defence equipment to over 85 nations. “Our target is to increase exports to Rs 35,000 crore in the next two years,” T Natarajan, Additional Secretary, Department of Defence Production was quoted as saying during June 2023. However, there are substantial hurdles in the way of India achieving its goal of Rs 35,000 crore worth of arms exports by 2025. For example, there is an over-reliance on exports of parts and components, rather than major defence equipment or the very few indigenously built platforms by India. Besides, some of Indian defence equipment exports, especially local military platforms like Dhruv Advanced Light helicopters (ALH) exported to Ecuador, were afflicted by problems of quality control as per reports.
A news outlet quoted the Indian defence ministry as saying that as of December 2022, expenditure on equipment from other nations was down 46 per cent compared to 2018-2019. CNBC reported that defence production in India during 2022-2023 has crossed Rs 1 lakh crore annually for the first time. Has India’s reliance on defence equipment imported from foreign nations reduced considerably?', 'uploads/Two_HAL_Tejas_flying_in_formation.jpeg', 'public'),
('ec201b9c-dedb-4d4b-9823-16bc517452b5', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '10 Steps To Become More Open-Minded', "Open-mindedness is the ability to be less judgmental and more inquisitive, introspective and considerate. Open-minded people may be viewed as more honest and reliable since they tend to consider multiple perspectives before reaching a decision. In this article, we discuss what it means to be open-minded, the benefits of having an open mindset and how you can work on building this ability.
What is open-mindedness?
Being open-minded means you are willing to look for and think about other perspectives. An open mindset is a belief that other people should be free to express their beliefs and arguments even if you do not agree with those views. Open-mindedness is often used as a synonym for being tolerant and not prejudiced. Open-minded people can fairly value experiences, beliefs, emotions, goals or arguments that may not align with their own. Open-minded people tend to be good listeners who strive to understand how other people perceive situations. They see differing opinions as valuable and thrive in diverse environments where they can interact with people with unique ideas and backgrounds.
Characteristics of open-mindedness
People who are open-minded share several characteristics such as:

    Willing to have their ideas challenged
    Staying calm when they are wrong
    Feeling empathy for other people
    Thinking about what others are thinking

Being humble

    about their knowledge and expertise
    Wanting to hear what other people have to say
    Believing everyone has a right to share their beliefs and thoughts

Benefits of being open-minded
Being open-minded offers mental and emotional benefits such as:

    Becoming more insightful about different beliefs, lifestyles and cultures in the world
    Gaining empathy for other people even if you have little in common
    Developing the ability to self-reflect and think more carefully about your own experiences, thoughts and decisions
    Boosting your confidence in your own beliefs and ideas
    Building mental strength and reasoning skills
    Improving your ability to be more realistic or optimistic
    Being more honest about your emotions with others and with yourself.

Related:6 Ways That Empathy Improves the Workplace
Get interview-ready with tips from Indeed
Prepare for interviews with practice questions and tips
How to be more open-minded
Becoming more open-minded takes time because changing the way your mind works can be challenging. Here are 10 ways to help you develop into a more open-minded person:
1. Be aware of your biases
Biases affect how we interpret information and can cause judgment or stereotypes. Becoming more aware of our own biases is the first step to challenging them and becoming open to new ideas.When you take in new information, consider how your biases affect your interpretation of that information. If you feel ready to accept the additional information, think about how this information might confirm your existing beliefs. If your response is to reject the information, consider what makes it challenging for you to accept.Related: 26 Logical Fallacies and How To Spot Them
2. Consider the opposite viewpoint
Open-mindedness involves being able to question not just others, but also yourself. Think about topics that you have strong opinions about and then imagine the perspective of someone on the other side of the argument. Try to think of at least three reasons someone might hold this opposing view. You might also consider reading a news article or book by or about someone who believes this opposing view. Being able to imagine another side's perspective, even if you continue to hold your previous beliefs, helps you to see the topic in a more nuanced way and become more open-minded when considering other subjects.
3. Practice generous interpretation
Generous interpretation is the process of assuming that people have good intentions. It goes beyond simply considering someone else's point of view and actively searches for positive justifications for their actions. For example, if a coworker does something that upsets you, think about the situation from their perspective. Consider what their motivation might be or what else might be happening in their life that affects their action. The most generous interpretation would assume the coworker meant well or was preoccupied with other things even if their actions were upsetting. It's OK to still acknowledge your negative feelings because of this coworker's actions while trying to practice sympathy for them.
4. Ask questions
Open-minded people tend to ask questions rather than offer their own opinions or argue. You can practice asking more questions during almost any conversation. For example, if a coworker is talking about a hobby you have never tried, ask them for more details about how it and what they enjoy about it.You can ask more questions about more emotionally charged topics as well. If a coworker starts talking about a political topic you have opposing views on, ask questions about their views rather than debate them. Seek to learn instead of immediately being defensive about your own opinions. If you decide to debate the coworker at a later date, you both can have a more complex conversation about the issue rather than simply arguing.Related: The Importance of Good Listening and How To Listen Effectively
5. Think about the neutral viewpoint
Try brainstorming the potential reasons for a neutral viewpoint on a controversial topic. You could also ask a friend if you could talk about this issue while they simply listen and summarize what you say. This technique of having someone else repeat your words back to you can help you see a situation more objectively.
6. Make new connections
Connecting with people who have a variety of perspectives on life is a good way to become more open-minded. Seek connections with people you may not often encounter or interact with for a prolonged period. For example, you could invite a coworker to lunch because you've never spent time with them outside of work before. Even if you have very little in common, take the opportunity to learn more about their interests and personality.
7. Experience different media
Considering other perspectives can become easier when we know more about the people who might have those perspectives. Reading books, watching movies, listening to podcasts or experiencing other types of media and art can all help us imagine or sympathize with different viewpoints.Look for media and art created by or about people with different views or experiences. For example, you can read a book written by someone from a different religion or watch a television show about people from another country.Related: Lifelong Learning: What Is It, Benefits and Habits
8. Join a new group
Changing how your mind works can be easier if you have a support system. One way to create this support system is to become involved with an open-minded group or organization.You might join a reading diversity challenge at your local library or bookstore. There may be a nonpartisan organization where people with various political views can come together to talk, organize and help others become informed about topics such as voting rights. If your supervisor approves, you could even start a group at your workplace.
9. Reframe negative thoughts
Open-minded people are typically more realistic or optimistic in their outlook. Try to be aware of when your mind immediately reaches a negative decision or conclusion. See if you can alter that negative thought and take a more neutral approach to a situation.For example, if you feel that a work task is impossible, try to think about the task as a challenge. Consider other obstacles you've overcome and what they have taught you. Acknowledge that it's OK to not fully succeed if you have done your personal best.Read more: 40 Tips To Overcome Being Negative
10. Acknowledge you are learning
Admitting when you are wrong can be challenging, but the ability to acknowledge your mistakes can help you become more open-minded. Realizing that your ideas could be incorrect can help you figure out how to learn from the experience and do better in the future. When someone helps you change your mind on a subject, show appreciation and thank them for putting the time and effort in to help you open your mind and educate yourself.", '', 'public'),
('cead0563-0215-427f-aca3-fea64be1130a', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Playing an instrument', 'Is anyone here playing an instrument and which one ?', '', 'public'),
('163fd188-e128-41f2-b3d7-0a3e95cb8dae', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'Welcome Word', 'Welcome to the group of news all around the world', 'uploads/475b59c61a6efb3ec8ac0a57e43fd03b.webp', 'public');


INSERT INTO groups (id, title, description, banner_url, creator_id) VALUES
('aa518af7-d623-4f5b-8bb6-bd7405862110', 'Tech Fans Club', 'This group is for the Tech Addicts to learn and share news in the Tech World...', 'uploads/TA-03032024-1-1024x576.webp', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232'),
('4c2c37cf-1d5a-4077-adc3-1687d852a570', 'Parlor', 'Group for open minded people to discuss about differents topics', 'uploads/LARBRE-A-PALABRES-620x395-1.png', 'c035df0d-8699-4880-a79e-1291915f70a9'),
('9c206815-391a-434b-8d36-15fb3df4dffd', 'World News', 'National and International news on politics, sports, business, tech ...', 'uploads/world_news_day.jpeg', '36db745b-c07c-492c-bfd8-aaec63aa6fd7');


INSERT INTO group_members (id, group_id, member_id, status, role) VALUES
('b6ca25c0-5bb7-43eb-b1dd-1be075afe6e2', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'accepted', 'admin'),
('26e64b28-0497-4f08-b3cf-db625a37b729', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'c035df0d-8699-4880-a79e-1291915f70a9', 'requesting', 'user'),
('f748b8d5-2f9e-4b5e-a272-9e47c3ecec73', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'accepted', 'admin'),
('56832149-f056-4f34-afb4-4b4ab134dbe9', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'accepted', 'user'),
('8fd09b61-1152-4f05-a6a5-2ec538ae394e', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'c035df0d-8699-4880-a79e-1291915f70a9', 'invited', 'user'),
('2f724f35-3648-4fea-84df-be12999ca83b', '9c206815-391a-434b-8d36-15fb3df4dffd', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'accepted', 'admin'),
('cc541aa3-99dd-48d1-acd4-864c6716abfd', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'declined', 'user');


INSERT INTO posts (id, group_id, user_id, title, content, image_url, privacy) VALUES
('067d1e18-0669-4222-baed-4c65259be792', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Discover the Latest in Sustainable Living', 'Join us as we explore eco-friendly products and tips to make a difference. #SustainableLiving #EcoFriendly', 'uploads/sustainable_living.jpeg', 'group'),
('b542446a-4c13-4198-943f-6ace73475ed0', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Healthy Recipes for a Healthier You', 'Dive into our collection of healthy recipes designed to boost your well-being. #HealthyRecipes #Wellness', 'uploads/fresh-fitness-food.avif', 'group'),
('6fa2cc57-94d6-4296-a0f3-89032ed68ccd', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Travel Tips for a Stress-Free Vacation', "Pack your bags and let's explore the best travel tips to ensure your vacation is as relaxing as it gets. #TravelTips #StressFree", '', 'group'),
('2381ed9b-17fe-4d10-989d-ac57182ff626', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Fashion Trends for the Season', 'Stay ahead of the curve with our latest fashion trends. #FashionTrends #SeasonalStyle', '', 'group'),
('3ffd68e8-eb57-4f73-8f67-e8f9df316c5a', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'Tech Gadgets That Will Change Your Life', 'Discover the latest tech gadgets that can make your life easier and more enjoyable. #TechGadgets #Innovation', 'uploads/449c71d8f06a9b5daba236cfd4a37a7e.jpg', 'group'),
('48da68ef-8a33-4255-9010-bf794a8cae6f', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'DIY Home Decor Ideas', 'Transform your space with these easy DIY home decor ideas. #DIY #HomeDecor', 'uploads/803eae848a2830191260cb36c2d29fd3.jpg', 'group'),
('153cb034-2c3e-4047-897b-eb55db0ded4a', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Books to Read This Summer', 'Escape into the world of books with our summer reading list. #SummerReads #Books', '', 'group'),
('8849c7c6-09cd-4f44-a2cc-87b0c3b65022', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Fitness Challenges for the New Year', 'Kickstart your fitness journey with our challenges designed to help you achieve your goals. #FitnessChallenges #NewYear', 'uploads/epic-cardio-challenge-intro.jpg', 'group'),
('d4ba832b-dbf3-4eae-a689-7d5ef80da0d6', '9c206815-391a-434b-8d36-15fb3df4dffd', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'Cultural Events to Attend', 'Discover the best cultural events happening around you. #CulturalEvents #Explore', '', 'group'),
('f2f6280f-4860-46a6-b0b8-5aff8ee88797', '9c206815-391a-434b-8d36-15fb3df4dffd', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'Eco-Friendly Products for Your Home', 'Make your home more sustainable with these eco-friendly products. #EcoFriendly #HomeSustainability', '', 'group');

INSERT INTO events (id, group_id, creator_id, title, description, date_time) VALUES
('6e07c132-6490-4c20-a5eb-87bec6717732', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Coding Session', 'This is a weekly coding session by "The Wizard"', '2024-05-05 00:00:00'),
('a4515c02-8a5e-450a-adf5-83562eeef13c', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'Members meeting', 'A zoom meeting call to welcome new members ...', '2024-03-31 00:00:00'),
('5bc4cc82-9713-4b9c-b31a-34505a218a7e', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'c035df0d-8699-4880-a79e-1291915f70a9', 'brainstorming', 'Room idea talk: quarter meeting for new ideas to make the group evolve', '2024-04-01 00:00:00');


INSERT INTO events_participants (id, event_id, member_id, response) VALUES
('0d1ad730-ef02-4ecb-9e4d-2c070183f26f', '6e07c132-6490-4c20-a5eb-87bec6717732', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'going'),
('077e4b39-94b9-4daa-ac67-4ffe6fb2950c', '6e07c132-6490-4c20-a5eb-87bec6717732', 'c035df0d-8699-4880-a79e-1291915f70a9', 'not_going'),
('90042b0a-2310-4931-b35d-594e244f3c9c', 'a4515c02-8a5e-450a-adf5-83562eeef13c', '498e640d-78d2-4171-b060-369d75c380ed', 'going'),
('c76d37b7-a46f-4fcf-8696-319c4fa41e01', 'a4515c02-8a5e-450a-adf5-83562eeef13c', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'not_going'),
('24108e9f-9dee-48fd-abe8-86c58b9a3ea1', '5bc4cc82-9713-4b9c-b31a-34505a218a7e', 'c035df0d-8699-4880-a79e-1291915f70a9', 'going'),
('85374be0-f0b4-4213-8b4f-2b445f7a2424', '5bc4cc82-9713-4b9c-b31a-34505a218a7e', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'not_going');

INSERT INTO private_messages (id, sender_id, receiver_id, content ) VALUES
('f115c834-09fe-4aba-8172-9e5034c48b34', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Hello Man, How is it?'), 
('b10fa83b-a742-4a9c-b04d-ad2dd84d5b39', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Hey, Im fine and you?'),
('48fcd310-8486-46a4-95fc-d95583430535', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Good, im blessed. Will you come at the match ?'),
('338206c3-5161-49f6-8714-966344adf150', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Not sure yet, maybe. Ill let you know.'),
('db6a8a3c-94af-4657-b5a0-6babf280a976', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Great, waiting for your confirmation.'),
('8800529f-3b8b-43fe-b47b-7ed14acfb150', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'okay'),
('318a186b-c8d8-48e0-a194-e7b6e58adc20', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Have to go, talk to you later'),
('1d68ef48-1edf-4dba-85cb-a836e744fdaa', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '498e640d-78d2-4171-b060-369d75c380ed', "what's up?"),
('7f0a1537-53e7-4c8f-af31-b7fb6635492f', '498e640d-78d2-4171-b060-369d75c380ed', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'yup boss, ready for the big day ?'),
('8022d94a-6a8a-4fe2-bb3d-64c791829eab', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '498e640d-78d2-4171-b060-369d75c380ed', 'of course... see ya there');


INSERT INTO group_messages (id, group_id, sender_id, content) VALUES
('aaae83f2-3095-4e1c-b537-73bc36c94ae1', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Welcome to the Group. Do a brief presentation and what make you joined the group.'),
('d365cd27-ff23-4c3f-ba4f-173d7ad35b22', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Hi im tester2 i think, and i am here as a tester for social network feature. lol'),
('dc660cca-ca64-4829-ba92-7eca94a9a966', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'This is the group official chat'),
('fabb94b7-87a4-464e-8913-c220130349d1', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'welcome onboard ...'),
('b19bcc7c-8288-4362-b729-e80903adb783', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Thank for the introcduction'),
('2896d78f-61ea-414c-9329-412971d0663e', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Glad to be here'),
('fbd251b8-4091-416d-9e28-cd915e215d08', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'Great guys, lets do it'),
('6fc22ba9-22ed-41d5-b697-4850715df1f0', '9c206815-391a-434b-8d36-15fb3df4dffd', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'Hello There ...');

INSERT INTO comments (id, user_id, post_id, content, image_url) VALUES
('98cb3c47-056a-4b15-8a89-8fc89c547aae', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '3e8958e6-b6ac-4eea-ab67-2fe811211978', 'Great Man, keep doing the good job!', 'uploads/default-avatar.png'),
('a018d89f-9a89-4f8e-8b42-8ad861e42664', 'c035df0d-8699-4880-a79e-1291915f70a9', 'aec2952f-fa1c-4065-9e23-f660819ac4ff', 'Great Man, keep doing the good job!', ''),
('e5aaa81b-33ae-4e8d-bf04-daf9687440ff', '498e640d-78d2-4171-b060-369d75c380ed', '3e8958e6-b6ac-4eea-ab67-2fe811211978', 'Great Man, keep doing the good job!', 'uploads/default-avatar.png'),
('8fd66b1a-53fd-4204-aeb4-dced6439ce71', '498e640d-78d2-4171-b060-369d75c380ed', 'cead0563-0215-427f-aca3-fea64be1130a', 'Great Man, keep doing the good job!', ''),
('c80f2e97-0368-42ba-a6ab-b850f041f1bc', '4f6f09ab-a290-4a38-9f70-9f61c2fd6e75', '163fd188-e128-41f2-b3d7-0a3e95cb8dae', 'Great Man, keep doing the good job!', 'uploads/default-avatar.png');

INSERT INTO selected_users (id, user_id, post_id) VALUES
('2d00d01a-1882-41a1-97c0-58492472ea87', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '9b77a08a-ff18-4022-97c7-4dc8f1a53230');

INSERT INTO notifications (id, user_id, concern_id, type) VALUES
('f5b43f4b-bdc6-4539-a0dd-d7ac218cc6a7', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'follow_request'),
('5255e83c-4bfe-4933-8fb6-781bae6fd335', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '498e640d-78d2-4171-b060-369d75c380ed', 'follow_accepted'),
('ce86e08a-4fdd-4a05-bfde-16acf40360c8', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'follow_declined'),
('0ebc6e97-98f7-47b0-8754-750f7efe5d73', '498e640d-78d2-4171-b060-369d75c380ed', 'c035df0d-8699-4880-a79e-1291915f70a9', 'follow_request'),
('52806c1c-b194-4bef-a9cb-ce988e45820b', '498e640d-78d2-4171-b060-369d75c380ed', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'follow_accepted'),
('db7a144b-c6c2-49c6-bdcf-cd5044a983fe', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'new_event'),
('d38a4335-fa20-4f32-b0b6-83db702ad747', '498e640d-78d2-4171-b060-369d75c380ed', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'new_event'),
('a018d89f-9a89-4f8e-8b42-8ad861e42664', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'new_event'),
('e5aaa81b-33ae-4e8d-bf04-daf9687440ff', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'c035df0d-8699-4880-a79e-1291915f70a9', 'new_message'),
('8fd66b1a-53fd-4204-aeb4-dced6439ce71', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'new_message'),
('c80f2e97-0368-42ba-a6ab-b850f041f1bc', 'c035df0d-8699-4880-a79e-1291915f70a9', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'new_message'),
('d28212f7-bb4b-48d5-a719-d786963affb2', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '498e640d-78d2-4171-b060-369d75c380ed', 'new_message'),
('517a1a87-15b1-4202-8711-c44aae95cf37', '498e640d-78d2-4171-b060-369d75c380ed', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'new_message'),
('2030d28a-8b3f-42ed-a8d8-b4d10445c683', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', '498e640d-78d2-4171-b060-369d75c380ed', 'new_message');
-- (),
-- (),
-- (),
-- (),
-- (),
-- ();
-- type TEXT CHECK(type = 'follow_request'OR type = 'follow_accepted' OR type = 'follow_declined' OR type = 'unfollow' OR type = 'group_invitation' OR type = 'new_message' OR type = 'new_event'),

-- INSERT INTO group_messages (id, group_id, sender_id, content) VALUES
-- ('aaae83f2-3095-4e1c-b537-73bc36c94ae1', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Welcome to the Group. Do a brief presentation and what make you joined the group.'),
-- ('d365cd27-ff23-4c3f-ba4f-173d7ad35b22', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Hi im tester2 i think, and i am here as a tester for social network feature. lol'),
-- ('dc660cca-ca64-4829-ba92-7eca94a9a966', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'This is the group official chat'),
-- ('fabb94b7-87a4-464e-8913-c220130349d1', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'welcome onboard ...'),
-- ('b19bcc7c-8288-4362-b729-e80903adb783', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'Thank for the introcduction'),
-- ('2896d78f-61ea-414c-9329-412971d0663e', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'c035df0d-8699-4880-a79e-1291915f70a9', 'Glad to be here'),
-- ('fbd251b8-4091-416d-9e28-cd915e215d08', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'Great guys, lets do it'),
-- ('6fc22ba9-22ed-41d5-b697-4850715df1f0', '9c206815-391a-434b-8d36-15fb3df4dffd', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'Hello There ...');

-- INSERT INTO group_members (id, group_id, member_id, status, role) VALUES
-- ('b6ca25c0-5bb7-43eb-b1dd-1be075afe6e2', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'accepted', 'admin'),
-- ('26e64b28-0497-4f08-b3cf-db625a37b729', 'aa518af7-d623-4f5b-8bb6-bd7405862110', 'c035df0d-8699-4880-a79e-1291915f70a9', 'requesting', 'user'),
-- ('f748b8d5-2f9e-4b5e-a272-9e47c3ecec73', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '498e640d-78d2-4171-b060-369d75c380ed', 'accepted', 'admin'),
-- ('56832149-f056-4f34-afb4-4b4ab134dbe9', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'accepted', 'user'),
-- ('8fd09b61-1152-4f05-a6a5-2ec538ae394e', '4c2c37cf-1d5a-4077-adc3-1687d852a570', 'c035df0d-8699-4880-a79e-1291915f70a9', 'invited', 'user'),
-- ('2f724f35-3648-4fea-84df-be12999ca83b', '9c206815-391a-434b-8d36-15fb3df4dffd', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'accepted', 'admin'),
-- ('cc541aa3-99dd-48d1-acd4-864c6716abfd', '4c2c37cf-1d5a-4077-adc3-1687d852a570', '36db745b-c07c-492c-bfd8-aaec63aa6fd7', 'declined', 'user');

-- INSERT INTO events_participants (id, event_id, member_id, response) VALUES
-- ('0d1ad730-ef02-4ecb-9e4d-2c070183f26f', '6e07c132-6490-4c20-a5eb-87bec6717732', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'going'),
-- ('077e4b39-94b9-4daa-ac67-4ffe6fb2950c', '6e07c132-6490-4c20-a5eb-87bec6717732', 'c035df0d-8699-4880-a79e-1291915f70a9', 'not_going'),
-- ('90042b0a-2310-4931-b35d-594e244f3c9c', 'a4515c02-8a5e-450a-adf5-83562eeef13c', '498e640d-78d2-4171-b060-369d75c380ed', 'going'),
-- ('c76d37b7-a46f-4fcf-8696-319c4fa41e01', 'a4515c02-8a5e-450a-adf5-83562eeef13c', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'not_going'),
-- ('24108e9f-9dee-48fd-abe8-86c58b9a3ea1', '5bc4cc82-9713-4b9c-b31a-34505a218a7e', 'c035df0d-8699-4880-a79e-1291915f70a9', 'going'),
-- ('85374be0-f0b4-4213-8b4f-2b445f7a2424', '5bc4cc82-9713-4b9c-b31a-34505a218a7e', 'a7ce8bfb-d026-4d5b-9c99-0d4c736c1232', 'not_going');