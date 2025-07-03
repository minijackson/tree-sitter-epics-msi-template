import XCTest
import SwiftTreeSitter
import TreeSitterEpicsMsiTemplate

final class TreeSitterEpicsMsiTemplateTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_epics_msi_template())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading EpicsMsiTemplate grammar")
    }
}
