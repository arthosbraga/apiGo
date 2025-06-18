S - Single Responsibility Principle (SRP)
The Single Responsibility Principle states that a class should have only one reason to change, meaning it should have only one job.

Incorrect Example: A single class handles both book data and its presentation logic.

C#
```cs
// Violates SRP
public class Book
{
    public string Title { get; set; }
    public string Author { get; set; }

    public void PrintToConsole()
    {
        Console.WriteLine($"Title: {Title}, Author: {Author}");
    }

    public string GetAsHtml()
    {
        return $"<h1>{Title}</h1><p>By {Author}</p>";
    }
}
```
Correct C# Example: We separate the concerns into two classes. The Book class only manages book data, and a separate BookPresenter class handles the presentation logic.

C#
```cs
// Adheres to SRP
public class Book
{
    public string Title { get; set; }
    public string Author { get; set; }
}

public class BookPresenter
{
    public void PrintToConsole(Book book)
    {
        Console.WriteLine($"Title: {book.Title}, Author: {book.Author}");
    }

    public string GetAsHtml(Book book)
    {
        return $"<h1>{book.Title}</h1><p>By {book.Author}</p>";
    }
}
```

O - Open/Closed Principle (OCP)
The Open/Closed Principle asserts that software entities should be open for extension but closed for modification. You should be able to add new functionality without changing existing code.

Incorrect Example: The PaymentProcessor class must be modified each time a new payment method is added.

C#
```cs
// Violates OCP
public enum PaymentType
{
    CreditCard,
    PayPal
}

public class PaymentProcessor
{
    public void ProcessPayment(PaymentType type)
    {
        if (type == PaymentType.CreditCard)
        {
            // Logic for credit card payment
            Console.WriteLine("Processing credit card payment...");
        }
        else if (type == PaymentType.PayPal)
        {
            // Logic for PayPal payment
            Console.WriteLine("Processing PayPal payment...");
        }
        // Adding a new payment type requires modifying this method.
    }
}
```

Correct C# Example: We use an interface. The PaymentProcessor works with any class that implements the IPaymentMethod interface. To add a new payment method, we create a new class, leaving the PaymentProcessor unchanged.

C#
```cs
// Adheres to OCP
public interface IPaymentMethod
{
    void ProcessPayment();
}

public class CreditCardPayment : IPaymentMethod
{
    public void ProcessPayment()
    {
        Console.WriteLine("Processing credit card payment...");
    }
}

public class PayPalPayment : IPaymentMethod
{
    public void ProcessPayment()
    {
        Console.WriteLine("Processing PayPal payment...");
    }
}

// We can easily add a new payment method
public class BankTransferPayment : IPaymentMethod
{
    public void ProcessPayment()
    {
        Console.WriteLine("Processing bank transfer payment...");
    }
}

public class PaymentProcessor
{
    public void ProcessPayment(IPaymentMethod paymentMethod)
    {
        paymentMethod.ProcessPayment();
    }
}
```

L - Liskov Substitution Principle (LSP)
The Liskov Substitution Principle states that objects of a superclass should be replaceable with objects of a subclass without affecting the program's correctness.

Incorrect Example: A Square inheriting from Rectangle can break expectations. A method that works with a Rectangle might not work correctly if a Square is passed to it.

C#
```cs
// Violates LSP
public class Rectangle
{
    public virtual int Width { get; set; }
    public virtual int Height { get; set; }

    public int GetArea()
    {
        return Width * Height;
    }
}

public class Square : Rectangle
{
    private int _side;

    public override int Width
    {
        get => _side;
        set
        {
            _side = value;
        }
    }

    public override int Height
    {
        get => _side;
        set
        {
            _side = value;
        }
    }
}

public class AreaCalculator
{
    public static void Calculate(Rectangle r)
    {
        r.Width = 5;
        r.Height = 10;
        // This assertion will fail for a Square, as setting Height also sets Width.
        Console.WriteLine($"Expected area: 50, Actual area: {r.GetArea()}");
    }
}

```
In the AreaCalculator, we expect to set the width and height independently. If we pass a Square object, setting the Height to 10 will also set the Width to 10, resulting in an area of 100, not the expected 50.

Correct C# Example: A better approach is to use a more generic interface like IShape that doesn't enforce a problematic inheritance hierarchy.

C#
```cs
// Adheres to LSP
public interface IShape
{
    int GetArea();
}

public class Rectangle : IShape
{
    public int Width { get; set; }
    public int Height { get; set; }

    public int GetArea()
    {
        return Width * Height;
    }
}

public class Square : IShape
{
    public int Side { get; set; }

    public int GetArea()
    {
        return Side * Side;
    }
}

```
I - Interface Segregation Principle (ISP)
The Interface Segregation Principle advises that clients should not be forced to depend on methods they do not use. It's better to have smaller, more specific interfaces.

Incorrect Example: A large IWorker interface forces a RobotWorker to implement an irrelevant Eat method.

C#
```cs
// Violates ISP
public interface IWorker
{
    void Work();
    void Eat();
}

public class HumanWorker : IWorker
{
    public void Work()
    {
        Console.WriteLine("Human working.");
    }

    public void Eat()
    {
        Console.WriteLine("Human eating.");
    }
}

public class RobotWorker : IWorker
{
    public void Work()
    {
        Console.WriteLine("Robot working.");
    }

    public void Eat()
    {
        // This method is irrelevant for a robot.
        throw new NotImplementedException("Robots don't eat!");
    }
}
```

Correct C# Example: We segregate the interface into smaller, more focused interfaces. Classes can then implement only the interfaces relevant to them.

C#
```cs
// Adheres to ISP
public interface IWorkable
{
    void Work();
}

public interface IEatable
{
    void Eat();
}

public class HumanWorker : IWorkable, IEatable
{
    public void Work()
    {
        Console.WriteLine("Human working.");
    }

    public void Eat()
    {
        Console.WriteLine("Human eating.");
    }
}

public class RobotWorker : IWorkable
{
    public void Work()
    {
        Console.WriteLine("Robot working.");
    }
}
```

D - Dependency Inversion Principle (DIP)
The Dependency Inversion Principle suggests that high-level modules should not depend on low-level modules; both should depend on abstractions.

Incorrect Example: The high-level Notification class directly depends on the low-level EmailClient class.

C#
```cs
// Violates DIP
public class EmailClient
{
    public void SendEmail()
    {
        Console.WriteLine("Sending email...");
    }
}

public class Notification
{
    private EmailClient _emailClient;

    public Notification()
    {
        _emailClient = new EmailClient(); // Tight coupling
    }

    public void Send()
    {
        _emailClient.SendEmail();
    }
}
```

Correct C# Example: We introduce an IMessageClient interface. The Notification class depends on this abstraction, and we can "inject" any implementation of the interface (like EmailClient or SmsClient). This decouples the high-level module from the low-level details.

C#
```cs
// Adheres to DIP
public interface IMessageClient
{
    void SendMessage();
}

public class EmailClient : IMessageClient
{
    public void SendMessage()
    {
        Console.WriteLine("Sending email...");
    }
}

public class SmsClient : IMessageClient
{
    public void SendMessage()
    {
        Console.WriteLine("Sending SMS...");
    }
}

// High-level module
public class Notification
{
    private readonly IMessageClient _messageClient;

    // The dependency is injected via the constructor.
    public Notification(IMessageClient messageClient)
    {
        _messageClient = messageClient;
    }

    public void Send()
    {
        _messageClient.SendMessage();
    }
}
```
```cs
// --- How to use it ---
// var emailNotifier = new Notification(new EmailClient());
// emailNotifier.Send();
//
// var smsNotifier = new Notification(new SmsClient());
// smsNotifier.Send();
```






