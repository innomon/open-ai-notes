# The Rectified Linear Unit (ReLU)

[https://share.google/aimode/CkFLn8SeB3pl4Kg4R](https://share.google/aimode/CkFLn8SeB3pl4Kg4R)

The Rectified Linear Unit, or ReLU, is the most commonly used activation function in deep learning. It is a piecewise linear function that outputs the input directly if it is positive; otherwise, it outputs zero. [1, 2, 3, 4, 5] 

Mathematically, it is defined as:
$$f(x) = \max(0, x)$$ 

## 1. Identify Function Behavior
The ReLU function acts as a gate. For any negative input, the neuron is "deactivated" (output is $0$). For any positive input, the output increases linearly with the input. This simplicity makes it computationally efficient compared to sigmoid or tanh functions. [6, 7, 8, 9, 10] 

## 2. Calculate the Derivative
The gradient of ReLU is essential for training via backpropagation:

* If $x > 0$, the derivative is $1$.
* If $x < 0$, the derivative is $0$.
* At $x = 0$, the derivative is technically undefined, but in practice, it is usually set to $0$ or $1$. [11, 12, 13, 14, 15] 

## 3. Evaluate Key Advantages

* **Sparse Activation**: In a random network, roughly half of the units are silent ($output = 0$), leading to sparse representations which are often more efficient.
* **Mitigates Vanishing Gradient**: Unlike sigmoid functions that "saturate" (flatten out) at high values, ReLU maintains a constant gradient of $1$ for all positive inputs, allowing deeper networks to train faster. [16, 17, 18, 19, 20] 

## 4. Recognize Potential Drawbacks

* **"Dying ReLU" Problem**: If a neuron receives a large negative bias, its output and gradient may stay at $0$ forever, effectively "dying" because it can no longer update its weights during training. Variations like Leaky ReLU ($f(x) = \max(0.01x, x)$) were created to fix this. [21, 22, 23, 24, 25] 

## Summary
ReLU is defined by $f(x) = \max(0, x)$. It is the standard choice for hidden layers in modern neural networks because it speeds up training and simplifies computation. [26, 27, 28, 29, 30] 

---

### References
[1] [https://en.wikipedia.org](https://en.wikipedia.org/wiki/Rectified_linear_unit#:~:text=Jarrett%20et%20al.%20%282009%29%20noted%20that%20rectification,neighboring%20filter%20outputs%20cancelling%20each%20other%20out.)  
[2] [https://www.kaggle.com](https://www.kaggle.com/code/dansbecker/rectified-linear-units-relu-in-deep-learning)  
[3] [https://deeplizard.com](https://deeplizard.com/lesson/dde2idlzra#:~:text=ReLU%20Activation%20Function%20%2D%20Deep%20Learning%20Dictionary,networks%2C%20like%20leaky%20and%20parametric%20%28%20%29.)  
[4] [https://towardsdatascience.com](https://towardsdatascience.com/what-are-activation-functions-in-deep-learning-cc4f01e1cf5c/#:~:text=This%20is%20a%20piecewise%20linear%20function%20that,is%20known%20a%20the%20%27rectified%27%20linear%20unit.)  
[5] [https://pub.towardsai.net](https://pub.towardsai.net/activation-functions-in-focus-understanding-relu-gelu-and-silu-841ed1c6df0c#:~:text=The%20Rectified%20Linear%20Unit%20%28ReLU%29%20is%20a,input%20if%20positive;%20otherwise%2C%20it%20outputs%20zero.)  
[6] [https://medium.com](https://medium.com/@nanda.yugandhar/beyond-relu-a-deep-dive-into-the-evolution-of-activation-functions-fcce7afc4221#:~:text=A%20powerful%20way%20to%20think%20about%20ReLU,hard%20switch%20with%20a%20smooth%2C%20data%2Ddependent%20%22dimmer.%22)  
[7] [https://www.dailydoseofds.com](https://www.dailydoseofds.com/why-is-relu-a-non-linear-activation-function/#:~:text=As%20depicted%20in%20the%20above%20equation%2C%20by,with%20a%20negative%20activation%20is%20turned%20off.)  
[8] [https://www.scribd.com](https://www.scribd.com/document/937580368/Activation-Functions-ReLU-Heuristics#:~:text=%E2%9D%96The%20main%20advantage%20of%20using%20the%20ReLU,means%20the%20neuron%20does%20not%20get%20activated.)  
[9] [https://www.thkp.co](https://www.thkp.co/blog/2019/4/14/nonlinearity-activation-functions#:~:text=One%20such%20non%2Dlinear%20activation%20function%20is%20called,effective%20activation%20functions%20for%20deep%20neural%20networks.)  
[10] [https://medium.com](https://medium.com/@prasanNH/activation-functions-in-neural-networks-b79a2608a106#:~:text=ReLU%20%28%20Rectified%20Linear%20Unit%20%29%20outputs,introduces%20a%20small%20slope/gradient%20for%20negative%20inputs.)  
[11] [https://www.studysmarter.co.uk](https://www.studysmarter.co.uk/explanations/engineering/artificial-intelligence-engineering/relu-function/#:~:text=Derivative%20of%20ReLU%20Function:%20The%20derivative%20is,f%27%28x%29%20=%200%20otherwise%2C%20facilitating%20gradient%2Dbased%20optimization.)  
[12] [https://medium.com](https://medium.com/@anushruthikae/all-about-activation-functions-choosing-the-right-activation-function-a63844e49a2a#:~:text=For%20all%20x%20values%20greater%20than%20zero%2C,is%20a%20constant%201%20for%20positive%20inputs.)  
[13] [https://medium.com](https://medium.com/@anushruthikae/all-about-activation-functions-choosing-the-right-activation-function-a63844e49a2a#:~:text=For%20all%20x%20values%20less%20than%20or,The%20derivative%20is%20zero%20for%20negative%20inputs.)  
[14] [https://dagshub.com](https://dagshub.com/glossary/rectified-linear-unit/#:~:text=The%20derivative%20of%20the%20ReLU%20function%20is,function%27s%20derivative%20to%20update%20the%20network%20weights.)  
[15] [https://iabac.org](https://iabac.org/blog/relu-activation-function#:~:text=ReLU%20has%20a%20%E2%80%9Ccorner%E2%80%9D%20at%20x%20=,no%20negative%20impact%20on%20training%20in%20practice.)  
[16] [https://en.wikipedia.org](https://en.wikipedia.org/wiki/Rectified_linear_unit#:~:text=Advantages%20Advantages%20of%20ReLU%20include:%20Sparse%20activation:,activation%20functions%20that%20saturate%20in%20both%20directions.)  
[17] [https://spotintelligence.com](https://spotintelligence.com/2023/06/16/activation-function/#:~:text=Sparsity:%20ReLU%20encourages%20sparsity%20in%20neural%20networks.,efficient%20and%20concise%20representations%20of%20the%20data.)  
[18] [https://medium.com](https://medium.com/analytics-vidhya/understanding-activation-functions-data-science-for-the-rest-of-us-b652048a064f#:~:text=They%20%28%20ReLU%20neurons%20%29%20%27re%20very,if%20you%20want%20to%20look%20it%20up%29.)  
[19] [https://iabac.org](https://iabac.org/blog/relu-activation-function#:~:text=ReLU%20avoids%20this%20by%20having%20a%20constant%20gradient%20of%201%20for%20all%20positive%20values.)  
[20] [https://eitca.org](https://eitca.org/artificial-intelligence/eitc-ai-adl-advanced-deep-learning/neural-networks/neural-networks-foundations/examination-review-neural-networks-foundations/what-are-the-key-differences-between-activation-functions-such-as-sigmoid-tanh-and-relu-and-how-do-they-impact-the-performance-and-training-of-neural-networks/)  
[21] [https://www.digitalocean.com](https://www.digitalocean.com/community/tutorials/relu-vs-elu-activation-function)  
[22] [https://www.educative.io](https://www.educative.io/answers/what-is-the-dying-relu-problem#:~:text=In%20neural%20networks%2C%20a%20biased%20term%20is,output%2C%20resulting%20in%20a%20dying%20ReLU%20problem.)  
[23] [https://www.analytixlabs.co.in](https://www.analytixlabs.co.in/blog/activation-function-in-neural-network/#:~:text=The%20issue%20with%20ReLU%20is%20a%20peculiar,network%27s%20ability%20to%20fit%20the%20data%20properly.)  
[24] [https://link.springer.com](https://link.springer.com/article/10.1007/s10044-024-01277-w#:~:text=It%20%28%20The%20ReLU%20function%20%29%20also,zero%20regardless%20of%20the%20input%20%5B%2012%5D.)  
[25] [https://www.dhiwise.com](https://www.dhiwise.com/post/activation-functions-neural-networks-guide#:~:text=Led%20to%20significant%20breakthroughs%20in%20deep%20learning,neuron%20effectively%20becomes%20%E2%80%9Cdead%E2%80%9D%20and%20stops%20learning.)  
[26] [https://www.researchgate.net](https://www.researchgate.net/figure/ReLU-defined-as-fx-max-0-x-is-a-representative-nonlinear-activation-function_fig3_381443573#:~:text=ReLU%2C%20defined%20as%20f%28x%29%20=%20max%20%280%2C,such%20as%20logistic%20sigmoid%20and%20hyperbolic%20tangent.)  
[27] [https://eureka.patsnap.com](https://eureka.patsnap.com/article/why-use-relu-instead-of-sigmoid-or-tanh#:~:text=Defined%20as%20f%28x%29%20=%20max%280%2C%20x%29%2C%20ReLU,ability%20to%20mitigate%20the%20vanishing%20gradient%20problem.)  
[28] [https://deepgram.com](https://deepgram.com/ai-glossary/activation-functions#:~:text=Otherwise%2C%20it%20%28%20ReLU%20%29%20outputs%20zero.,affecting%20the%20receptive%20fields%20of%20convolutional%20layers.)  
[29] [https://medium.com](https://medium.com/@chawthirisan/why-is-non-linearity-important-in-neural-networks-and-activation-functions-101-992af6e819db#:~:text=ReLU%20became%20popular%20because%20it%20reduces%20the,default%20activation%20function%20in%20modern%20deep%20learning.)  
[30] [https://medium.com](https://medium.com/@pinaki.brahma/a-primer-in-activation-functions-ef925ff2384a#:~:text=Hidden%20Layers%20in%20Deep%20Neural%20Networks:%20ReLU,computationally%20inexpensive%2C%20facilitating%20faster%20training%20and%20convergence.)
