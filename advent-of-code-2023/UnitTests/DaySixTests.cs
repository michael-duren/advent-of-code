using Shared;
using Shared.Solutions.DaySix;

namespace UnitTests
{
    public class DaySixTests
    {
        private readonly List<string> _testInput;
        private readonly List<string> _input;

        public DaySixTests()
        {
            string testPath = PathInputFactory.Create("Six").Test;
            string inputPath = PathInputFactory.Create("Six").Input;

            _testInput = Helpers.ReadInput(testPath);
            _input = Helpers.ReadInput(inputPath);
        }

        [Fact]
        public void PartTwoTest()
        {
            // arrange
            const long expected = 71503;
            // act 
            long result = PartTwo.Solve(_testInput);

            // assert
            Assert.Equal(expected, result);
        }
    }
}