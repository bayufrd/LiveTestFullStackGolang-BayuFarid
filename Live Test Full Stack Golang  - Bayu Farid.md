## **Full Stack Go Lang Live Coding Test**

**Instructions for Candidate:**

●     Please read each question carefully.

●     Write your Go code solution in the shared document provided.

●     The use of AI is not permitted

●     Ensure your code is complete and runnable for each question, including `package main`, necessary `import` statements, and a `func main()` to demonstrate your solution where applicable.

●     You have approximately 30 minutes to complete all questions.

---

###


### **Question 1: Structs and Methods (Approx. 7 minutes)**

Define a struct in Go called `Product` with two fields:

1. `Name` (string)  
2. `Price` (float64)

Then, write a method for this `Product` struct called `GetPriceWithTax`. This method should take a `taxRate` (float64, e.g., 0.1 for 10% tax) as an argument and return the price of the product *including* the tax.

Finally, show how you would create an instance of `Product` with Name "Laptop" and Price 1200.0, and then call `GetPriceWithTax` with a tax rate of 0.07 (7%) and print the result to the console.

---

Your Answer :

###  

### **Question 2: Functions, Slices, and Error Handling (Approx. 8 minutes)**

Write a Go function called `CalculateAverage` that:

1. Takes a slice of integers (`[]int`) as input.  
2. Returns two values:

   ○     The average of the numbers in the slice as a `float64`.

   ○     An `error`.

3. If the input slice is empty, the function should return `0.0` for the average and an error message "cannot calculate average of an empty slice".  
4. Otherwise, it should calculate and return the average and `nil` for the error.

Show how you would call this function within `func main()`:

●     First, with a sample slice `[]int{10, 20, 30, 40}`. Print the average if there's no error, or print the error if one occurs.  
●     Second, with an empty slice. Print the error if one occurs.

---

Your Answer :

###  

### **Question 3: Maps and Iteration (Approx. 8 minutes)**

Declare a map in Go called `inventory` where keys are product names (strings) and values are their quantities (integers). Initialize this map with the following data:

●     `"apple"`: 10

●     `"banana"`: 5

●     `"orange"`: 0

Write Go code (within `func main()`) to iterate through this `inventory` map. For each product:

●     If the quantity is 0, print: `"[ProductName] is out of stock."` (e.g., "orange is out of stock.")  
●     Otherwise, print: `"[ProductName] has [Quantity] items."` (e.g., "apple has 10 items.")

---

Your Answer :

## **ReactJS & TypeScript Live Coding Test: Simple Tag Input 🏷️**

**Welcome\!** We're excited to see your skills in action. This test is designed to assess your ability to work with ReactJS and TypeScript in a short, focused exercise.

**Duration:** 30 minutes

**Objective:** The goal of this test is to build a simple "Tag Input" component. We want to see how you approach creating interactive UI elements, manage state, handle user events, and utilize TypeScript for type safety.

---

### **Problem Statement:**

You are tasked with creating a **Simple Tag Input** component. This component should allow users to:

1. **Enter Text:** Provide an input field where users can type tag names.  
2. **Add Tags:** When the user presses the "Enter" key after typing in the input field, the entered text should be added as a new tag.  
3. **Display Tags:** Show the added tags visually on the screen (e.g., as small styled boxes or "pills" above or near the input field).  
4. **Remove Tags:** Allow users to remove an already added tag by clicking on a designated area for that tag (e.g., a small "x" icon or button next to each tag).  
5. **Validation:**  
   * Prevent empty strings from being added as tags.  
   * Prevent duplicate tags from being added.

---

### **Requirements & Acceptance Criteria:**

* **Framework/Libraries:**  
  * Use **React functional components** and **Hooks** (primarily `useState`).  
  * Utilize **TypeScript** for defining types for props (if any), state, and function signatures.  
  * **No external libraries** for state management (like Redux) or UI components are needed. Focus on core React and TypeScript.  
* **Functionality:**  
  * An input field is present for typing tags.  
  * Tags are added to a displayed list when the "Enter" key is pressed in the input field.  
  * The input field should clear after a tag is successfully added.  
  * Each displayed tag should have a way to be removed (e.g., a clickable "x").  
  * The component should not allow adding empty tags (e.g., a string with only spaces, or an empty string).  
  * The component should not allow adding a tag if it already exists in the list (case-sensitive or case-insensitive, please clarify with your interviewer if unsure, but a case-sensitive check is fine for this exercise).  
* **Styling:**  
  * Basic styling is sufficient. You can use inline styles or simple CSS. The primary focus is on functionality and React/TypeScript implementation, not on making it look perfect, though adding basic interactive styles like hover, focus, and active states would be a great bonus.  
* **Provided Setup:**  
  * You should initialize your project using \`npx create-react-app tag-input \--template typescript\` or a similar command to set up a basic React TypeScript environment.

---

### **What We're Looking For:**

* **React Fundamentals:** Your understanding of components, state, props, JSX, and event handling.  
* **TypeScript Proficiency:** How you use types for state, props, and functions to ensure type safety.  
* **Problem Solving:** Your approach to breaking down the problem and implementing the features.  
* **Code Clarity:** Writing readable and maintainable code.  
* **Communication:** Write out any assumptions you made.

**Optional Objectives:**

* **Testing Ability:** Demonstrating the ability to write tests for your components.  
* **Atomic Design Principles:** Applying atomic design principles in creating your components.  
* **Design Ability:** Showing your design acumen in the component's visual presentation (basic styling still applies, but extra design considerations will be noted).  
* **Usability:** Considering user experience and making the component intuitive to use.

---

### **Tips for Success:**

* **Iterative Approach:** Start with the core functionality (e.g., displaying tags, then adding, then removing) and build up from there.  
* **Focus on Functionality Over Perfect Styling:** Getting the logic right is more important than pixel-perfect CSS in this exercise.  
* **Don't Panic if You Get Stuck:** It's okay. Take a moment, think it through, state your assumptions  
* **Manage Your Time:** Keep an eye on the 30-minute time limit. It's okay if not every single detail is polished; a working core is great.

Good luck\! We look forward to seeing what you build.

Your Answer :

